package utils

import (
	logger "alto_server/common/log"
	"alto_server/constants"
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func DirExists(dirPath string) (bool, error) {
	// 获取文件信息
	info, err := os.Stat(dirPath)
	if err != nil {
		// 如果错误是“文件不存在”，则返回 false 和 nil 错误
		if os.IsNotExist(err) {
			return false, nil
		}
		// 其他错误（如权限不足）返回错误
		return false, err
	}
	// 存在且是目录
	return info.IsDir(), nil
}

// UncompressFiles Extracts a zip file to ./xsl/extracted/<zipname>/.
// If filePath is "path/to/downloaded/file.zip" or empty, the first zip file found in ./xsl is used.
// On success, extractedMigrationRoot is set to the extracted root directory.
func UncompressFiles(zipPath string, destRoot string) error {
	exists, err := DirExists(destRoot)
	if err != nil {
		return err
	}
	if !exists {
		if err := os.MkdirAll(destRoot, 0o755); err != nil {
			return fmt.Errorf("Failed to create the extraction directory: %w", err)
		}
	} else {
		return nil
	}
	// If files do not exist, uncompress the zip file
	fmt.Printf("UncompressFiles: Extracting %s\n", zipPath)
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		fmt.Printf("UncompressFiles: Error opening zip file: %v\n", err)
		return fmt.Errorf("Failed to open the zip file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		target := filepath.Join(destRoot, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(target, f.Mode()); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		in, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			in.Close()
			return err
		}
		if _, err := io.Copy(out, in); err != nil {
			in.Close()
			out.Close()
			return err
		}
		in.Close()
		out.Close()
	}

	fmt.Printf("UncompressFiles: Extracted to %s\n", destRoot)
	return nil
}

// DF16 -> DF
func GetOltSeriesByType(neType string) string {
	return neType[0:2]
}
func GetOLTPortIndexByCard2(card string) (string, error) {
	// if card == "" {
	// 	return "", errors.New("invalid card name")
	// }
	if strings.Compare(card, "nt") == 0 {
		return constants.NT_PORT, nil
	}
	if strings.Compare(card, "ihub") == 0 {
		return constants.IHUB_PORT, nil
	}
	if strings.Compare(card, "df") == 0 {
		return constants.DF_PORT, nil
	}
	if strings.Contains(card, "lt") {
		index, err := strconv.Atoi(card[2:])
		if err != nil {
			return "", errors.New("invalid card name")
		}

		return strconv.Itoa(constants.LT_PORT_START + index), nil
	}

	return "", errors.New("invalid card name")
}
func GetOLTCardByPort(port string) (string, error) {
	// if card == "" {
	// 	return "", errors.New("invalid card name")
	// }
	if strings.Compare(port, "832") == 0 {
		return "nt", nil
	} else if strings.Compare(port, "831") == 0 {
		return "ihub", nil
	} else if strings.Compare(port, "833") == 0 {
		return "lt1", nil
	} else if strings.Compare(port, "834") == 0 {
		return "lt2", nil
	}
	return "", errors.New("invalid card name")
}
func ProcessRPCFile(inputCfg string, outputPath string, modLogger *logrus.Logger) error {
	// fmt.Printf("ProcessRPCFile: normalizing %s -> %s\n", inputPath, outputPath)
	// b, err := ioutil.ReadFile(inputPath)
	// if err != nil {
	// return fmt.Errorf("failed to read RPC file: %w", err)
	// }
	// txt := string(b)
	reXML := regexp.MustCompile(`(?m)^\s*<\?xml[^>]*>\s*`)
	txt := reXML.ReplaceAllString(inputCfg, "")
	reRPCOpen := regexp.MustCompile(`(?is)<rpc-reply[^>]*>`)
	reRPCClose := regexp.MustCompile(`(?is)</rpc-reply>`)
	txt = reRPCOpen.ReplaceAllString(txt, "")
	txt = reRPCClose.ReplaceAllString(txt, "")
	reDataOpen := regexp.MustCompile(`(?is)<data[^>]*>`)
	reDataClose := regexp.MustCompile(`(?is)</data>`)
	txt = reDataOpen.ReplaceAllString(txt, "")
	txt = reDataClose.ReplaceAllString(txt, "")
	txt = strings.TrimSpace(txt)
	wrapped := fmt.Sprintf(`<config xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">%s</config>`, txt)
	if err := os.WriteFile(outputPath, []byte(wrapped), 0o644); err != nil {
		modLogger.Debug("ailed to write normalized file  ", err)
		return fmt.Errorf("failed to write normalized file: %w", err)
	}
	fmt.Println("ProcessRPCFile: normalization complete")
	modLogger.Debug("ProcessRPCFile: normalization complete  ")
	return nil
}

func ExecuteXSLTransformations(inputFile string, outputFile string, fromV string, targetV string, modLogger *logrus.Logger, extractedMigrationRoot string) error {
	if inputFile == "" || outputFile == "" {
		return errors.New("input and output required")
	}
	// fmt.Printf("ExecuteXSLTransformations: input=%s output=%s\n", inputFile, outputFile)
	/*
		extractedMigrationRoot := "common/migrationTools/xsl/extracted/device-extension-ls-df-cfxr-e-25.6-298"
		if strings.Compare(port, "831") == 0 {
			extractedMigrationRoot = "common/migrationTools/xsl/extracted/ls_mf_ihub_lmnt-b_mf2_lmxr-b_25.9_459"
		} else if strings.Compare(port, "832") == 0 {
			extractedMigrationRoot = "common/migrationTools/xsl/extracted/ls_mf_lmnt-b_mf2_lmxr-b_25.9_459"
		} else if strings.Compare(port, "830") == 0 {
			extractedMigrationRoot = "common/migrationTools/xsl/extracted/ls_df_cfxr-e_25.6_298"
		} else {
			extractedMigrationRoot = "common/migrationTools/xsl/extracted/ls_mf_lwlt-c_25.9_459"
		}*/
	// ensure migration data extracted
	// if extractedMigrationRoot == "" {
	// 	// try to find migration under ./xsl/extracted or directly ./xsl/migration
	// 	candidates := []string{
	// 		filepath.Join(".", "xsl", "migration"),
	// 		filepath.Join(".", "xsl", "extracted"),
	// 	}
	// 	for _, c := range candidates {
	// 		if fi, err := os.Stat(c); err == nil && fi.IsDir() {
	// 			// try to find migration.xml under c
	// 			mx := filepath.Join(c, "migration.xml")
	// 			if _, err := os.Stat(mx); err == nil {
	// 				extractedMigrationRoot = c
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// logger.SystemLogger.Error("ExecuteXSLTransformations   extractedMigrationRoot: ", extractedMigrationRoot)
	logger.SystemLogger.Errorf("ExecuteXSLTransformations   extractedMigrationRoot: %s\n", extractedMigrationRoot)
	// normalize RPC input to a temp initial file
	initial := inputFile
	// if file doesn't start with <config ...>, process it
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}
	if !strings.Contains(string(content), "<config") {
		tmp := filepath.Join(".", "convert_db_0_input.xml")
		if err := ProcessRPCFile(inputFile, tmp, modLogger); err != nil {
			return fmt.Errorf("ProcessRPCFile failed: %w", err)
		}
		initial = tmp
	}

	// --- NEW: run a single pre-transformation (sim_av_mig.xsl 或 PRE_XSL) BEFORE chain execution ---
	preXsl := GetMigrationToolPrefixPath() + "sim_av_mig.xsl"
	// if preXsl == "" {
	// 	// prefer extracted migration's sim_av_mig.xsl
	// 	if extractedMigrationRoot != "" {
	// 		cand := filepath.Join(extractedMigrationRoot, "xsl", "sim_av_mig.xsl")
	// 		if _, err := os.Stat(cand); err == nil {
	// 			preXsl = cand
	// 		}
	// 	}
	// 	// fallback to project ./xsl/sim_av_mig.xsl
	// 	if preXsl == "" {
	// 		cand2 := filepath.Join(".", "xsl", "sim_av_mig.xsl")
	// 		if _, err := os.Stat(cand2); err == nil {
	// 			preXsl = cand2
	// 		}
	// 	}
	// }
	// logger.SystemLogger.Error("ExecuteXSLTransformations   preXsl: ", preXsl)
	if preXsl != "" {
		preOut := strings.TrimSuffix(initial, filepath.Ext(initial)) + "_pre.xml"
		logger.SystemLogger.Errorf("ExecuteXSLTransformations: running pre-transformation %s -> %s\n", preXsl, preOut)
		if err := RunSaxon(preXsl, initial, preOut); err != nil {
			return fmt.Errorf("pre-transformation failed: %w", err)
		}
		initial = preOut
		logger.SystemLogger.Errorf("ExecuteXSLTransformations: pre-transformation complete, new initial=%s\n", initial)
	} else {
		logger.SystemLogger.Errorf("ExecuteXSLTransformations: no pre-transformation found (PRE_XSL or sim_av_mig.xsl)")
	}
	// --- END NEW ---

	logger.SystemLogger.Errorf("ExecuteXSLTransformations: SOURCE=%s TARGET=%s extractedMigrationRoot=%s\n", fromV, targetV, extractedMigrationRoot)

	// if SOURCE/TARGET provided and migration.xml available -> chain execution
	if fromV != "" && targetV != "" && extractedMigrationRoot != "" {
		logger.SystemLogger.Error("start generate new file:", outputFile)
		migXML := filepath.Join(extractedMigrationRoot, "migration.xml")
		if _, err := os.Stat(migXML); err != nil {
			// try parent
			migXML = filepath.Join(extractedMigrationRoot, "migration", "migration.xml")
		}
		if _, err := os.Stat(migXML); err == nil {
			logger.SystemLogger.Errorf("ExecuteXSLTransformations: using migration.xml: %s\n", migXML)
			migrations, err := ParseMigrationXML(migXML)
			if err != nil {
				logger.SystemLogger.Error(err.Error())
				return fmt.Errorf("failed to parse migration.xml: %w", err)
			}
			chain, err := BuildScriptChain(fromV, targetV, migrations)
			if err != nil {
				logger.SystemLogger.Error(err.Error())
				return fmt.Errorf("failed to build script chain: %w", err)
			}
			logger.SystemLogger.Errorf("ExecuteXSLTransformations: script chain length=%d\n", len(chain))
			for i, s := range chain {
				fmt.Printf("  chain[%d] = %s\n", i, s)
			}
			if len(chain) == 0 {
				// nothing to do, copy initial to outputFile
				inputBytes, _ := ioutil.ReadFile(initial)
				if err := ioutil.WriteFile(outputFile, inputBytes, 0o644); err != nil {
					logger.SystemLogger.Error(err.Error())
					return fmt.Errorf("failed to write final output: %w", err)
				}
				logger.SystemLogger.Errorf("ExecuteXSLTransformations: no scripts to run, copied input to output")
				return nil
			}
			// ensure xslDir: migration/xsl under extracted root, or ./xsl/migration/xsl
			xslDir := filepath.Join(extractedMigrationRoot, "xsl")
			if fi, err := os.Stat(xslDir); err != nil || !fi.IsDir() {
				// try extractedMigrationRoot itself if it already is the migration dir
				xslDir = filepath.Join(extractedMigrationRoot, "migration", "xsl")
			}
			logger.SystemLogger.Errorf("ExecuteXSLTransformations: xslDir=%s\n", xslDir)
			// execute chain producing intermediate files convert_db_0.xml ...
			final, err := ExecuteMigrationChain(chain, xslDir, initial, strings.TrimSuffix(outputFile, filepath.Ext(outputFile)), modLogger)
			if err != nil {
				logger.SystemLogger.Error(err.Error())
				return err
			}
			// move final to requested outputFile if necessary
			if final != outputFile {
				data, err := ioutil.ReadFile(final)
				if err != nil {
					logger.SystemLogger.Error(err.Error())
					return fmt.Errorf("failed to read final file: %w", err)
				}
				if err := ioutil.WriteFile(outputFile, data, 0o644); err != nil {
					logger.SystemLogger.Error(err.Error())
					return fmt.Errorf("failed to write %s: %w", outputFile, err)
				}
			}
			logger.SystemLogger.Errorf("ExecuteXSLTransformations: chain execution complete")
			return nil
		} else {
			logger.SystemLogger.Error("not do real work")
		}
	} else {
		logger.SystemLogger.Errorf("input error:%s %s %s ", fromV, targetV, extractedMigrationRoot)
	}

	// otherwise run single XSL (CURRENT_XSL or ./xsl/sim_av_mig.xsl)
	// xsl := ""
	// if xsl == "" {
	// 	// prefer migration's sim_av_mig.xsl if exists
	// 	if extractedMigrationRoot != "" {
	// 		cand := filepath.Join(extractedMigrationRoot, "xsl", "sim_av_mig.xsl")
	// 		if _, err := os.Stat(cand); err == nil {
	// 			xsl = cand
	// 		}
	// 	}
	// 	if xsl == "" {
	// 		cand2 := filepath.Join(".", "xsl", "sim_av_mig.xsl")
	// 		if _, err := os.Stat(cand2); err == nil {
	// 			xsl = cand2
	// 		}
	// 	}
	// 	if xsl == "" {
	// 		fmt.Println("ExecuteXSLTransformations: sim_av_mig.xsl not found")
	// 		return errors.New("no XSL specified and sim_av_mig.xsl not found (set CURRENT_XSL or place sim_av_mig.xsl in ./xsl)")
	// 	}
	// }
	// fmt.Printf("ExecuteXSLTransformations: running single xsl: %s\n", xsl)
	// if err := RunSaxon(xsl, initial, outputFile); err != nil {
	// 	return err
	// }
	logger.SystemLogger.Errorf("ExecuteXSLTransformations: single xsl complete")
	return nil
}

type migrationFile struct {
	XMLName    xml.Name    `xml:"migrations"`
	Migrations []migration `xml:"migration"`
}

type migration struct {
	Source  string  `xml:"source"`
	Target  string  `xml:"target"`
	Scripts scripts `xml:"scripts"`
}

type scripts struct {
	Script []script `xml:"script"`
}

type script struct {
	Subtree string `xml:"subtree,attr"`
	Name    string `xml:",chardata"`
}

// ParseMigrationXML parses migration.xml and returns a list of migration entries.
func ParseMigrationXML(path string) ([]migration, error) {
	if path == "" {
		return nil, errors.New("migration xml path required")
	}
	fmt.Printf("ParseMigrationXML: parsing %s\n", path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read migration xml: %w", err)
	}
	var mf migrationFile
	if err := xml.Unmarshal(b, &mf); err != nil {
		return nil, fmt.Errorf("unmarshal migration xml: %w", err)
	}
	return mf.Migrations, nil
}

// BuildScriptChain builds an ordered list of script filenames to run to migrate from source to target.
func BuildScriptChain(source, target string, migrations []migration) ([]string, error) {
	fmt.Printf("BuildScriptChain: source: %s\n", source)
	fmt.Printf("BuildScriptChain: target: %s\n", target)
	if source == "" || target == "" {
		return nil, errors.New("source and target required")
	}
	if strings.TrimSpace(source) == strings.TrimSpace(target) {
		return nil, nil
	}
	mmap := make(map[string]migration)
	for _, m := range migrations {
		mmap[strings.TrimSpace(m.Source)] = m
	}
	var chain []string
	cur := strings.TrimSpace(source)
	tgt := strings.TrimSpace(target)
	visited := make(map[string]bool)
	for {
		if cur == tgt {
			break
		}
		if visited[cur] {
			return nil, fmt.Errorf("cycle detected in migration chain at %s", cur)
		}
		visited[cur] = true
		m, ok := mmap[cur]
		if !ok {
			return nil, fmt.Errorf("no migration entry for %s (cannot reach %s)", cur, tgt)
		}
		for _, s := range m.Scripts.Script {
			name := strings.TrimSpace(s.Name)
			if name != "" {
				chain = append(chain, name)
			}
		}
		cur = strings.TrimSpace(m.Target)
		if cur == "" {
			return nil, fmt.Errorf("migration entry for %s has empty target", m.Source)
		}
	}
	if cur != tgt {
		return nil, fmt.Errorf("could not reach target %s (stopped at %s)", tgt, cur)
	}
	fmt.Printf("BuildScriptChain: built chain from %s to %s: %v\n", source, target, chain)
	return chain, nil
}

// RunSaxon 执行单个 XSL 脚本。
func RunSaxon(xslPath, inputFile, outputFile string) error {
	if xslPath == "" || inputFile == "" || outputFile == "" {
		return errors.New("xsl, input and output required")
	}
	javaExe, saxonJar, err := getJavaAndSaxon()
	if err != nil {
		return err
	}
	fmt.Printf("RunSaxon: java=%s saxon=%s xsl=%s input=%s output=%s\n", javaExe, saxonJar, xslPath, inputFile, outputFile)
	args := []string{"-jar", saxonJar, fmt.Sprintf("-o:%s", outputFile), fmt.Sprintf("-xsl:%s", xslPath), fmt.Sprintf("-s:%s", inputFile)}
	fmt.Printf("RunSaxon: running command: %s %s\n", javaExe, strings.Join(args, " "))
	cmd := exec.Command(javaExe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("saxon execution failed: %w", err)
	}
	fmt.Println("RunSaxon: completed")
	return nil
}

// ExecuteMigrationChain 按顺序执行脚本链，输出文件以 outPrefix_i.xml 命名并返回最终文件路径。
func ExecuteMigrationChain(chain []string, xslDir, inputFile, outPrefix string, modLogger *logrus.Logger) (string, error) {
	if len(chain) == 0 {
		return inputFile, nil
	}
	if xslDir == "" {
		xslDir = filepath.Join(".", "xsl")
	}
	fmt.Printf("ExecuteMigrationChain: xslDir=%s input=%s outPrefix=%s\n", xslDir, inputFile, outPrefix)
	curIn := inputFile
	for i, sc := range chain {
		xslPath := sc
		if !filepath.IsAbs(sc) {
			xslPath = filepath.Join(xslDir, sc)
		}
		outFile := fmt.Sprintf("%s_%d.xml", outPrefix, i)
		fmt.Printf("ExecuteMigrationChain: step %d, xsl=%s -> out=%s\n", i, xslPath, outFile)
		if err := RunSaxon(xslPath, curIn, outFile); err != nil {
			return "", fmt.Errorf("failed at step %d script %s: %w", i, sc, err)
		}
		curIn = outFile
	}
	fmt.Printf("ExecuteMigrationChain: final output=%s\n", curIn)
	return curIn, nil
}

func getJavaAndSaxon() (string, string, error) {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)

	javaCandidates := []string{
		filepath.Join(exeDir, "migrationTools", "jre", "bin", "java.exe"),
		filepath.Join(".", "migrationTools", "jre", "bin", "java.exe"),
		filepath.Join(exeDir, "assets", "jre", "windows-x64", "bin", "java.exe"),
	}
	javaExe := ""
	for _, c := range javaCandidates {
		if _, err := os.Stat(c); err == nil {
			javaExe = c
			break
		}
	}
	if javaExe == "" {
		if p, err := exec.LookPath("java"); err == nil {
			javaExe = p
		}
	}
	if javaExe == "" {
		fmt.Println("getJavaAndSaxon: java not found in candidates or PATH")
		return "", "", errors.New("java not found: place embedded JRE in migrationTools/jre or install java in PATH")
	}

	saxonCandidates := []string{
		filepath.Join(exeDir, "migrationTools", "saxon9he.jar"),
		filepath.Join(".", "migrationTools", "saxon9he.jar"),
		filepath.Join(exeDir, "assets", "migrationTools", "saxon9he.jar"),
	}
	saxon := ""
	for _, s := range saxonCandidates {
		if _, err := os.Stat(s); err == nil {
			saxon = s
			break
		}
	}
	if saxon == "" {
		fmt.Println("getJavaAndSaxon: saxon9he.jar not found in candidates")
		return "", "", errors.New("saxon9he.jar not found: place saxon9he.jar in third_party/ or assets/third_party/")
	}

	fmt.Printf("getJavaAndSaxon: selected java=%s saxon=%s\n", javaExe, saxon)
	return javaExe, saxon, nil
}

// GenerateConfigurationPackages reads FINAL_XML, NE_TYPE, VERSION (or TARGET as fallback), OUT_ZIP
// environment variables and creates a zip package containing the final xml and a small metadata yml.
// If FINAL_XML isn't set, defaults to convert_db_2.xml.
func GenerateConfigurationPackages(outputFile string, neType string, srcVer string, targetVer string) error {

	if _, err := os.Stat(outputFile); err != nil {
		fmt.Printf("GenerateConfigurationPackages: final xml not found: %s\n", outputFile)
		return fmt.Errorf("final xml not found: %w", err)
	}

	outZip := "./software/convert_package.zip"
	logger.SystemLogger.Errorf("GenerateConfigurationPackages: xml=%s ne_type=%s version=%s out_zip=%s\n", outputFile, neType, targetVer, outZip)

	src := srcVer
	dst := targetVer
	cleanedNeType := CleanNeType(neType)
	// ---- New logic: Automatically generate a name with SOURCE, TARGET, and NE_TYPE when outZip is the default value ----
	if filepath.Base(outZip) == "convert_package.zip" {
		safeSrc := strings.ReplaceAll(src, "/", "-")
		safeDst := strings.ReplaceAll(dst, "/", "-")
		safeNe := strings.ReplaceAll(cleanedNeType, "/", "-")
		outZip = fmt.Sprintf("%s-%s_%s_new_db.zip", safeSrc, safeDst, safeNe)
		fmt.Printf("GenerateConfigurationPackages: Auto-renamed output to %s\n", outZip)
	}

	fmt.Printf("GenerateConfigurationPackages: XML=%s NE_TYPE=%s VERSION=%s OUT_ZIP=%s\n", outputFile, neType, targetVer, outZip)

	// ---- Prepare a temporary YML file first (for adding/overwriting the YML in the ZIP) ----
	ymlContent := fmt.Sprintf("ne_type: %s\nversion: %s\nhash:11000111\n", neType, targetVer)
	tmpYmlName := strings.ToLower(neType) + ".yml" // YML filename (e.g., ne123.yml)
	tmpYmlPath := fmt.Sprintf("%s/%s", os.TempDir(), tmpYmlName)
	// Create the temporary YML file
	tmpYmlFile, err := os.Create(tmpYmlPath)
	if err != nil {
		return fmt.Errorf("Failed to create the temporary YML file: %w", err)
	}
	if _, err := tmpYmlFile.WriteString(ymlContent); err != nil {
		tmpYmlFile.Close()
		os.Remove(tmpYmlPath)
		return fmt.Errorf("Failed to write content to the YML file: %w", err)
	}
	tmpYmlFile.Close()
	defer os.Remove(tmpYmlPath) // Delete the temporary file when the function exits

	// ---- Define the target filenames to be added/overwritten ----
	targetXmlName := filepath.Base(outputFile) // Target XML name (e.g., convert_db_2.xml)
	targetYmlName := tmpYmlName                // Target YML name (e.g., ne123.yml)

	// ---- Scenario 1: ZIP file already exists → Copy existing files + Add/overwrite XML/YML ----
	if _, err := os.Stat(outZip); err == nil {
		fmt.Printf("GenerateConfigurationPackages: %s exists; adding/overwriting XML and YML...\n", outZip)
		// 1. Create a temporary ZIP file (to avoid modifying the original file directly)
		tmpZipPath := outZip + ".tmp"
		newZipFile, err := os.Create(tmpZipPath)
		if err != nil {
			return fmt.Errorf("Failed to create the temporary ZIP file: %w", err)
		}
		newZipWriter := zip.NewWriter(newZipFile)

		// 2. Open the old ZIP file and copy all existing files (skip XML/YML with the same name; add them uniformly later)
		oldZipReader, err := zip.OpenReader(outZip)
		if err != nil {
			newZipWriter.Close()
			newZipFile.Close()
			return fmt.Errorf("Failed to open the old ZIP file: %w", err)
		}
		for _, oldFile := range oldZipReader.File {
			// Skip files with the same name as the target XML/YML (they will be overwritten with new files later)
			if oldFile.Name == targetXmlName || oldFile.Name == targetYmlName {
				fmt.Printf("GenerateConfigurationPackages: Will overwrite the old file: %s\n", oldFile.Name)
				continue
			}
			// Copy other files to the temporary ZIP file as-is
			newFileInZip, err := newZipWriter.Create(oldFile.Name)
			if err != nil {
				oldZipReader.Close()
				newZipWriter.Close()
				newZipFile.Close()
				return fmt.Errorf("Failed to create a file in the temporary ZIP file: %w", err)
			}
			oldFileReader, err := oldFile.Open()
			if err != nil {
				oldZipReader.Close()
				newZipWriter.Close()
				newZipFile.Close()
				return fmt.Errorf("Failed to open the old file: %w", err)
			}
			if _, err := io.Copy(newFileInZip, oldFileReader); err != nil {
				oldFileReader.Close()
				oldZipReader.Close()
				newZipWriter.Close()
				newZipFile.Close()
				return fmt.Errorf("Failed to copy the old file: %w", err)
			}
			oldFileReader.Close()
		}
		oldZipReader.Close() // Close the old ZIP reader

		// 3. Add/overwrite XML and YML to the temporary ZIP file
		if err := addFileToZip(newZipWriter, outputFile, targetXmlName); err != nil {
			newZipWriter.Close()
			newZipFile.Close()
			return fmt.Errorf("Failed to add XML to the temporary ZIP file: %w", err)
		}
		if err := addFileToZip(newZipWriter, tmpYmlPath, targetYmlName); err != nil {
			newZipWriter.Close()
			newZipFile.Close()
			return fmt.Errorf("Failed to add YML to the temporary ZIP file: %w", err)
		}

		// 4. Complete writing to the temporary ZIP file and replace the original ZIP file
		if err := newZipWriter.Close(); err != nil {
			newZipFile.Close()
			return fmt.Errorf("Failed to close the temporary ZIP writer: %w", err)
		}
		newZipFile.Close()
		if err := os.Rename(tmpZipPath, outZip); err != nil {
			return fmt.Errorf("Failed to replace the old ZIP file: %w", err)
		}

		fmt.Printf("GenerateConfigurationPackages: Added/overwrote XML and YML in %s\n", outZip)
		return nil
	}

	// ---- Scenario 2: ZIP file does not exist → Create a new ZIP file and add XML and YML ----
	newZipFile, err := os.Create(outZip)
	if err != nil {
		return fmt.Errorf("Failed to create the new ZIP file: %w", err)
	}
	defer newZipFile.Close()

	newZipWriter := zip.NewWriter(newZipFile)
	// Add XML
	if err := addFileToZip(newZipWriter, outputFile, targetXmlName); err != nil {
		newZipWriter.Close()
		return fmt.Errorf("Failed to add XML to the new ZIP file: %w", err)
	}
	// Add YML
	if err := addFileToZip(newZipWriter, tmpYmlPath, targetYmlName); err != nil {
		newZipWriter.Close()
		return fmt.Errorf("Failed to add YML to the new ZIP file: %w", err)
	}

	if err := newZipWriter.Close(); err != nil {
		return fmt.Errorf("Failed to close the new ZIP writer: %w", err)
	}
	fmt.Printf("GenerateConfigurationPackages: Created %s with XML + YML\n", outZip)
	return nil
}
func GenerateZip(inFileArray []string, outZip string) error {
	newZipFile, err := os.Create(outZip)
	if err != nil {
		return fmt.Errorf("Failed to create the new ZIP file: %w", err)
	}
	defer newZipFile.Close()

	newZipWriter := zip.NewWriter(newZipFile)
	for _, inFile := range inFileArray {
		if err := addFileToZip(newZipWriter, inFile, filepath.Base(inFile)); err != nil {
			newZipWriter.Close()
			return fmt.Errorf("Failed to add file to the new ZIP file: %w", err)
		}
	}
	if err := newZipWriter.Close(); err != nil {
		return fmt.Errorf("Failed to close the new ZIP writer: %w", err)
	}
	logger.SystemLogger.Error("GenerateConfigurationPackages: Created %s with XML + YML\n", outZip)
	return nil
}
func addFileToZip(zw *zip.Writer, src, name string) error {
	fmt.Printf("GenerateConfigurationPackages: Adding %s as %s\n", src, name)
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := zw.Create(name)
	if err != nil {
		return err
	}

	// If XML declaration deletion is required, retain this section; otherwise, use io.Copy(f, w) directly
	if strings.HasSuffix(strings.ToLower(name), ".xml") {
		content, err := io.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Failed to read XML: %w", err)
		}
		// Precisely delete the target XML declaration
		targetDecl1 := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n")
		targetDecl2 := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
		targetDecl3 := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		switch {
		case bytes.HasPrefix(content, targetDecl1):
			content = content[len(targetDecl1):]
		case bytes.HasPrefix(content, targetDecl2):
			content = content[len(targetDecl2):]
		case bytes.HasPrefix(content, targetDecl3):
			content = content[len(targetDecl3):]
		}
		if _, err := w.Write(content); err != nil {
			return fmt.Errorf("Failed to write XML: %w", err)
		}
		return nil
	}

	_, err = io.Copy(w, f)
	return err
}

func CleanNeType(neType string) string {
	if neType == "" {
		return ""
	}
	parts := strings.Split(neType, "-")
	filtered := make([]string, 0, len(parts))
	for _, part := range parts {
		if strings.ToUpper(part) != "IHUB" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, "-")
}

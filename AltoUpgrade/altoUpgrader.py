#!/usr/bin/python3
import requests
import subprocess
import json
import sys
import os
from  common import ReadConfig
from  common import ErrorPrint

cfg = ReadConfig()

url = cfg.get_ini("info","url")
print(url)
dstPort = cfg.get_ini("info","dstPort")
print(dstPort)
swVersion = cfg.get_ini("info","swVersion")
print(swVersion)
oltIp = cfg.get_ini("info","oltIp")
print(oltIp)
swName = "L6GQAG"+swVersion

def checkConnection(oltIp, dstPort):
    process = subprocess.Popen('/usr/bin/nc -vz '+oltIp + ' ' +dstPort, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output, error = process.communicate()
    if "succeeded" in error.decode('utf-8'):
        return True
    else:
        return False

def download(dstPort, swVersion):
    dirName = 'lightspan_'+swVersion
    swName = "L6GQAG"+swVersion
    downloadUrl = dirName+"/L6GQAG/"+swName
    payload = {'dstPort': dstPort,
    'oltId': '192.168.1.1',
    'action':'download',
    'url': downloadUrl,
    'name': swName}
    print(payload)
    ErrorPrint(json.dumps(payload))
    
    response = requests.request("PUT", url, data=payload)
    
    print(response.text)
    ErrorPrint(response.text)
    res = (json.loads(response.text))
    
    xmlinfo = (res['data'])
    errorInfo = (xmlinfo['Errors'])
    print(errorInfo)
    if errorInfo == None:
        print("download ok, no error")
        ErrorPrint("download ok, no error")
def active(dstPort, swVersion):
    swName = "L6GQAG"+swVersion
    payload = {'dstPort': dstPort,
    'oltId': '192.168.1.1',
    'name': swName,
    'action': 'active'}
    files=[
    
    ]
    headers = {
    
    }
    print(payload)
    ErrorPrint(json.dumps(payload))
    response = requests.request("PUT", url, headers=headers, data=payload, files=files)
    print(response.text)
    ErrorPrint(response.text)
    res = (json.loads(response.text))
    
    xmlinfo = (res['data'])
    errorInfo = (xmlinfo['Errors'])
    print(errorInfo)
    if errorInfo == None:
        print("active ok, no error")
        ErrorPrint("active ok, no error")
def commit(dstPort, swVersion):
    swName = "L6GQAG"+swVersion
    payload = {'dstPort': dstPort,
    'oltId': '192.168.1.1',
    'name': swName,
    'action': 'commit'}
    files=[
    
    ]
    headers = {
    
    }
    print(payload)
    ErrorPrint(json.dumps(payload))
    response = requests.request("PUT", url, headers=headers, data=payload, files=files)
    print(response.text)
    ErrorPrint(response.text)
    res = (json.loads(response.text))
    
    xmlinfo = (res['data'])
    errorInfo = (xmlinfo['Errors'])
    print(errorInfo)
    if errorInfo == None:
        print("commit ok, no error")
        ErrorPrint("commit ok, no error")

def checkProcessNum():
    process = subprocess.Popen('ps -ef|grep altoUpgrader|grep -v grep|wc -l', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output, error = process.communicate()
    return (output.decode('utf-8').strip('\n'))

#check 1 daemon
os.system("/etc/init.d/networking restart")
currProcessNum = checkProcessNum()
print(currProcessNum)
ErrorPrint(currProcessNum)
if currProcessNum == '2':
    print("just one, could continue")
elif currProcessNum == '1':
    print("just one, could continue")
else:
    print("more than one, need exit")
    sys.exit()

connStatus = checkConnection(oltIp, dstPort)
if connStatus == True:
    print("connection to olt ok,could action")
    ErrorPrint("connection to olt ok,could take action")
else:
    print("connection to olt nok, could not action")
    ErrorPrint("connection to olt nok, could not take action")
    sys.exit()



querystring = {"oltId":"192.168.1.1","dstPort":dstPort}
response = requests.request("GET", url, params=querystring)
res = (json.loads(response.text))
#print(res['olt_software_info']['HardwareState']['Component']['Software']['Software']['Revisions']['Revision'])
xmlinfo = json.loads(res['olt_software_info'])
oltversionInfo = (xmlinfo['HardwareState']['Component']['Software']['Software']['Revisions']['Revision'])
downloadState = (xmlinfo['HardwareState']['Component']['Software']['Software']['Download']['CurrentState']['State'])
oltVersionLen  = len(oltversionInfo)
downloaded = False
if downloadState == "idle":
    for i in range(oltVersionLen):
        if swName == oltversionInfo[i]["Name"]:
            print(oltversionInfo[i])
            ErrorPrint(json.dumps(oltversionInfo[i]))
            if oltversionInfo[i]["IsValid"] == 'true':
                downloaded = True
                if oltversionInfo[i]["IsActive"] == 'false':
                    print(swName + "download done, could active it")
                    ErrorPrint(swName + "download done, could active it")
                    os.system("/usr/bin/mplayer /root/music/active.mp3 &")
                    active(dstPort,swVersion)
                else:
                    if oltversionInfo[i]["IsCommitted"] == 'false':
                        print(swName + "active done, could commit it")
                        ErrorPrint(swName + "active done, could commit it")
                        os.system("/usr/bin/mplayer /root/music/commit.mp3 &")
                        commit(dstPort,swVersion)
                    else:
                        print(swName + "active done,  commit done, no need action")
                        ErrorPrint(swName + "active done,  commit done, no need action")
                        os.system("/usr/bin/mplayer /root/music/upgradeDone.mp3 &")

    if downloaded == False:
        print("need download")
        ErrorPrint("need download")
        download(dstPort,swVersion)
else:
    print("downloading, skip")
    ErrorPrint("downloading, skip")

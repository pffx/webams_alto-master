#!/usr/bin/python3
import requests
import subprocess
import json
import sys
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


querystring = {"oltId":"192.168.1.1","dstPort":dstPort}
response = requests.request("GET", url, params=querystring)
res = (json.loads(response.text))
print(res['olt_software_info'])
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
                else:
                    if oltversionInfo[i]["IsCommitted"] == 'false':
                        print(swName + "active done, could commit it")
                        ErrorPrint(swName + "active done, could commit it")
                    else:
                        print(swName + "active done,  commit done, no need action")

    if downloaded == False:
        print("need download")
        ErrorPrint("need download")
else:
    print("downloading, skip")
    ErrorPrint("downloading, skip")

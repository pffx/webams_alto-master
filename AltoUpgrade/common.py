import base64
import os
import configparser
import time

proDir = os.path.split(os.path.realpath(__file__))[0]
configPath = os.path.join(proDir, "config.ini")
print(configPath)

def ErrorPrint(message):
    localtime = time.asctime( time.localtime(time.time()))
    logfile=open("/root/altoUpgrade.log", "a");
    logfile.write("["+localtime+"]");
    logfile.write(message);
    logfile.write("\n");
    logfile.close()

class ReadConfig:
    def __init__(self):
        self.cf = configparser.ConfigParser()
        self.cf.read(configPath)
    def get_ini(self, par, name):
        value = self.cf.get(par, name)
        return value

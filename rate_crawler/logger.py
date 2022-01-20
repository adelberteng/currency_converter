import logging
import json
import os
from datetime import datetime

import requests

logger_level_tbl = {
    "debug": logging.DEBUG,
    "info": logging.INFO,
    "warning": logging.WARNING,
    "error": logging.ERROR,
    "critical": logging.CRITICAL
}

level_val = {
    "debug": 10,
    "info": 20,
    "warning": 30,
    "error": 40,
    "critical": 50
}

_top_dir_path = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
log_dir_path = os.path.join(_top_dir_path, "log")

class Logger:
    """ A self-made logger base on Python build-in logger.  
    This logger will create a logger with Slack alert, 
    and set logger_level or log output easily.

    :param logger_level: set a level to filter which messages 
        will be recorded. Default is info.
    :param log_folder: Default is None, which means do not record log to
        log file in local folder. If this argument is given, logger will
        record the log to the specified folder and named by date.
    :param console_enable: Default is True, which means show the log 
        on console. if that is not you want, put False in this parameter.
    :param slack_webhook_url: If you want to print alert messages to slack, 
        put your slack webhook url in this parameter.
    :param content_type: slack header content, you can pass this.
    :param env: The environ parameter will show where the message come from.
    """
    # Singleton
    instance = None
    def __new__(cls, *args, **kwargs):
        if cls.instance is None:
            cls.instance = super().__new__(cls)
        return cls.instance

    def __init__(
        self, 
        logger_level = "info", 
        log_folder = None, 
        console_enable = True,
        slack_webhook_url = None, 
        content_type = "application/json",
        env = "DEV"
        ):
        self.logger_level = logger_level
        self.log_folder = log_folder
        self.console_enable = console_enable
        self.slack_webhook_url = slack_webhook_url
        self.header = {'Content-type': content_type}
        self.env = env.upper()
        self.logger = self._create_logger(
            self.logger_level, self.log_folder, console_enable = self.console_enable)

    @staticmethod
    def _create_logger(logger_level = "info", log_folder = None, console_enable = True):
        logging.captureWarnings(True)
        formatter = logging.Formatter("[%(levelname)s] %(asctime)s %(message)s")
        logger = logging.getLogger(str(log_folder))
        logger.setLevel(logger_level_tbl[logger_level])

        if log_folder:
            if not os.path.exists(log_dir_path + "/" + log_folder):
                os.makedirs(log_dir_path + "/" + log_folder)

            filename = "{:%Y-%m-%d}".format(datetime.now())
            fileHandler = logging.FileHandler(
                f"{log_dir_path}/{log_folder}/{filename}.log", "a", "utf-8")
            fileHandler.setFormatter(formatter)
            logger.addHandler(fileHandler)

        if console_enable is True:
            consoleHandler = logging.StreamHandler()
            consoleHandler.setLevel(logging.DEBUG)
            consoleHandler.setFormatter(formatter)
            logger.addHandler(consoleHandler)

        return logger


    def debug(self, msg):
        self.logger.debug(str(msg))
        if (self.slack_webhook_url is not None 
            and level_val["debug"] >= level_val[self.logger_level]):
            msg_json = json.dumps({"text" : f"{self.env}-DEBUG: {str(msg)}"})
            requests.post(self.slack_webhook_url, headers= self.header, data = msg_json)

    def info(self, msg):
        self.logger.info(str(msg))
        if (self.slack_webhook_url is not None 
            and level_val["info"] >= level_val[self.logger_level]):
            msg_json = json.dumps({"text" : f"{self.env}-INFO: {str(msg)}"})
            requests.post(self.slack_webhook_url, headers= self.header, data = msg_json)
    
    def warning(self, msg):
        self.logger.warning(str(msg))
        if (self.slack_webhook_url is not None 
            and level_val["warning"] >= level_val[self.logger_level]):
            msg_json = json.dumps({"text" : f"{self.env}-WARNING: {str(msg)}"})
            requests.post(self.slack_webhook_url, headers= self.header, data = msg_json)
    
    def error(self, msg):
        self.logger.error(str(msg))
        if (self.slack_webhook_url is not None 
            and level_val["error"] >= level_val[self.logger_level]):
            msg_json = json.dumps({"text" : f"{self.env}-ERROR: {str(msg)}"})
            requests.post(self.slack_webhook_url, headers= self.header, data = msg_json)
    
    def critical(self, msg):
        self.logger.critical(str(msg))
        if (self.slack_webhook_url is not None 
            and level_val["critical"] >= level_val[self.logger_level]):
            msg_json = json.dumps({"text" : f"{self.env}-CRITICAL: {str(msg)}"})
            requests.post(self.slack_webhook_url, headers= self.header, data = msg_json)

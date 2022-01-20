import sys
import time

import requests
from bs4 import BeautifulSoup

class RateCrawler:
	def __init__(self):
		self.url = "https://rate.bot.com.tw/xrt?Lang=en-US"
		self.headers = {
			"User-agent": (
				"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0)"
				"AppleWebKit/537.36 (KHTML, like Gecko) "
				"Chrome/86.0.4240.198 Safari/537.36"
			),
		}

	@staticmethod
	def _get_rate_from_tr(tr):
		cols = tr.find_all("td")
		try:
			for col in cols:
				if col.attrs["data-table"] == "Cash Buying":
					return float(list(col)[0])
		except Exception as e:
			sys.exit(e)

	def get_rate_dict(self):
		res = requests.get(self.url, headers=self.headers)
		soup = BeautifulSoup(res.text, "html.parser")
		tbody = soup.find("tbody")
		rows = tbody.find_all("tr")

		rate_dict = {"TWD": 1.0}
		for tr in rows:
			if tr.find_all("div", class_="sp-america-div"):
				rate_dict["USD"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-japan-div"):
				rate_dict["JPY"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-korea-div"):
				rate_dict["KRW"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-hong-kong-div"):
				rate_dict["HKD"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-singapore-div"):
				rate_dict["SGD"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-euro-div"):
				rate_dict["EUR"] = self._get_rate_from_tr(tr)
			elif tr.find_all("div", class_="sp-china-div"):
				rate_dict["CNY"] = self._get_rate_from_tr(tr)

		return rate_dict

	@staticmethod
	def mapping_rate(rate_dict):
		mapping_rate_dict = {}
		for src_c in rate_dict.keys():
			exchange_dict = {}
			for target_c in rate_dict.keys():
				exchange_dict[target_c] =  round(rate_dict[src_c]/rate_dict[target_c], 3)
			mapping_rate_dict[src_c] = exchange_dict

		return mapping_rate_dict


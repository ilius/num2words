import gzip
import sys
import unittest
from os.path import abspath, dirname

sys.path.insert(0, dirname(abspath(__file__)))
from en import convert_int, convert_string


def loadTestData() -> list[tuple[str, int, str]]:
	data: list[tuple[str, int, str]] = []
	with gzip.open("test-data.gz", mode="rt") as _file:
		for line in _file:
			line = line.strip()
			if not line:
				continue
			parts = line.split("\t")
			num_str = parts[0]
			num = int(num_str)
			words = parts[1]
			data.append((num_str, num, words))
	return data


class TestConvertByTestData(unittest.TestCase):
	def __init__(self, *args, **kwargs):
		unittest.TestCase.__init__(self, *args, **kwargs)
		self.testData = loadTestData()

	def test_convert_string(self):
		for num_str, _num, words in self.testData:
			self.assertEqual(convert_string(num_str), words)

	def test_convert_int(self):
		for _num_str, num, words in self.testData:
			self.assertEqual(convert_int(num), words)


if __name__ == "__main__":
	unittest.main()

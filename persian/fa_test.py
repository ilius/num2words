import gzip
import sys
import unittest
from os.path import abspath, dirname

sys.path.insert(0, dirname(abspath(__file__)))
from fa import convert_int, convert_int_ordinal, convert_str, convert_str_ordinal


def loadTestData(fname: str) -> list[tuple[str, int, str]]:
	data: list[tuple[str, int, str]] = []
	with gzip.open(fname, mode="rt") as _file:
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
		self.testData = loadTestData("test-data.gz")
		self.ordinalTestData = loadTestData("test-data-ordinal.gz")

	def test_convert_string(self):
		for num_str, _num, words in self.testData:
			self.assertEqual(convert_str(num_str), words)

	def test_convert_int(self):
		for _num_str, num, words in self.testData:
			self.assertEqual(convert_int(num), words)

	def test_convert_string_ordinal(self):
		for num_str, _num, words in self.ordinalTestData:
			self.assertEqual(convert_str_ordinal(num_str), words)

	def test_convert_int_ordinal(self):
		for _num_str, num, words in self.ordinalTestData:
			self.assertEqual(convert_int_ordinal(num), words)


if __name__ == "__main__":
	unittest.main()

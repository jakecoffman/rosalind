import dna
import unittest

class Test(unittest.TestCase):
    def test_dna(self):
        d = 'AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC'
        actual = dna.count_nucleotides(d)
        self.assertEqual([20, 12, 17, 21], actual)

if __name__ == "__main__":
    unittest.main()
    

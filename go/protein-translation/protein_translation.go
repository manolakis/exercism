package protein

import "errors"

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
	noProteins     []string
)

func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])

		if errors.Is(err, ErrInvalidBase) {
			return noProteins, err
		}

		if errors.Is(err, ErrStop) {
			break
		}

		proteins = append(proteins, protein)
	}

	return proteins, err
}

func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return protein, err
}

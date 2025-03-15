package educative0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   int
	}{
		{
			name:   "test 1",
			input1: "wxyz",
			input2: "zwxgy",
			want:   3,
		},
		{
			name:   "test 2",
			input1: "cbda",
			input2: "abc",
			want:   2,
		},
		{
			name:   "test 3",
			input1: "aaaaa",
			input2: "aaaaaa",
			want:   0,
		},
		{
			name:   "test 4",
			input1: "courae",
			input2: "couearg",
			want:   6,
		},
		{
			name:   "test 5",
			input1: "hello",
			input2: "helo",
			want:   2,
		},
		{
			name:   "test 6",
			input1: "oxtpjwhyclgtmwktcgabyyrykwgpurwyxpqtpnkxpzedjduccwioknospsonbudzvofxzymtrrqfqaduwyhhtbgrdeckpiestiivmfxwvrfcvbovlkvpbzcfpiumfgxkwjnsmrqyelnvmjjeveyfeexacbanjtqcugbfohgzscnuurdcziedaummfdtlcxhxxdjcxghzvidlcyfvkvamsbqflukrnkkaalkwsxtzowjrjiwhjmxtgherrxvnlwenxjvtggaviszcbgtlenguzeqhrqmyjrvdjcjyzagkrgxymcaetwzhhwlownvwzvettnufkjysdqvyroenrnfzhkcyzvafacqyimzdiqwssqaaecwtkikhcebwwmrdvalvthrrxckxbowjppaqdkkvrzxeagglhtkgxpkmjxcbrduymokkkvkojnnmnfstjpvhhjavycysvhkmmmcssxqfwxroxojvxcjqyncpdofbdrpzbsvpoeuwmmbdlfzsbfpzpygvopnuautrdvnllslcgpvvsglcnordahwsfjtcbgbqpitgalsenzpwrpjpljgpnigsnckdfjkhsknibjnzodyjusvzudclaypqfbatmsviywmxvifhbnwvdoddocdawfvmzwcbnegucwedbnydcyvljejzmiodlgspvyhskjyopndujoxokedfszedwsasvaieeyjivoczljmbgujvawqhzrofadtvkshhzzasmmaockhhukzfvhkxxovtcohlqhmuhtgfmbdzhdqvoknrpupjseyqxmwstgtlndxlpoyfqaqecywduvqlyrriwlifcnavbyvvavltudejoflcqmrwezzzbncwjzddwvagawzstklknkowlznfhgghcpcbyrinqghnvbaheywsttgjrpbwejsqjocrbqatfqotrekiudmbauaowbrqtfgnxdnntxefjmgskixruscyxkuezmsvtkvskygvjrfosyo",
			input2: "ozpjvzhhabpsaooaxaorgmmmakpkspkqnatfosnyxktuvofclhmyyosoalypdtyrltbqcasmorkhdvqbrlckxnythzrmyzdgmncfcycrulkhmjdngehujwgbpaobunzpsowjjrspezuhkfaazkeqnxtplwzcxsdrrfyucecimvdtvvpvtwaidcvlrbbwqxrtipkaxidhhrlkecfsdqcwjdrbtkeabbbjojlrvnlizsbcywpsrfgiymuzyabqwvkwegcfewnvfvzfgknhjwntwwruzrzkfcynxjublmgcolpkpcrtlmkpjehsmdbcjkrxctvyeveozracekbmoaavcdyfxskwtaneqggeuxyszjvzfplhuqjywepanufdnmaevldyxyvfmasedfdsbhvgyvmproukkgesdzjjovbxabgivqenpvqxzzfnkzdoevgwrvmbfrmzvlawkxdwwwfevmwodkkvjwmjosaplxogvjmhcyjoeldlzhxtgforstxknjjnjludatcwdddtfsodzdfvmhhcykqgiwihhdsegimmansgqnatqakdghypbpxgylvsfqjyzojtkvvunzcenggxgbohednvkyosctaqovrswtpaircruleztwrsvwhqhdctecmvxvauhfuavydsgstdnizboblsxvtcvzzaqpwhnnvjdezjgxodqdayrrkrlngtprowwrzxacpwnhncygoucvhfspcipjzdvnuwccjwhpqpokfswxqttixgniymkjmnikunveqljvkoyzxzdsudonlqwybcjulignjxqgxzvcqwejswhesgebvkttpcfmsnvctsltavqfmmuiwnexkddoihjvzfzysbxhrbrsyufoadtukhyezryxklzxjpbkjsjsrnidnjomfhthvaflgitefabdqcjwmcufwnihqfmywobccuuqvvlvtqgybgwuelrxetfjikkpnbijxcjmvrgmccsgkysiokwygw",
			want:   117,
		},
	}

	for _, test := range tests {
		have := extraCharacterIndex(test.input1, test.input2)
		assert.Equalf(t, test.want, have, "%s: extraCharacterIndex not "+"calculated correctly", test.name)
	}
}

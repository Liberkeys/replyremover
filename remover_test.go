package replyremover

import (
	"bufio"
	"io/ioutil"
	"strings"
	"testing"
)

const FixturesPath = "fixtures/"
const FixturesExtension = ".txt"

const CommonBody = `Fusce bibendum, quam hendrerit sagittis tempor, dui turpis tempus erat, pharetra sodales ante sem sit amet metus.
Nulla malesuada, orci non vulputate lobortis, massa felis pharetra ex, convallis consectetur ex libero eget ante.
Nam vel turpis posuere, rhoncus ligula in, venenatis orci. Duis interdum venenatis ex a rutrum.
Duis ut libero eu lectus consequat consequat ut vel lorem. Vestibulum convallis lectus urna,
et mollis ligula rutrum quis. Fusce sed odio id arcu varius aliquet nec nec nibh.`

func getFixture(name string, t *testing.T) string {
	b, err := ioutil.ReadFile(FixturesPath + name + FixturesExtension)
	if err != nil {
		t.Fatalf("Failed to load fixture '%s%s%s'", FixturesPath, name, FixturesExtension)
	}
	return string(b)
}

func compare(str1, str2 string, t *testing.T) {
	if str1 != str2 {
		scanner := bufio.NewScanner(strings.NewReader(str2))
		for scanner.Scan() {
			t.Log(scanner.Text())
		}
		t.Error("Strings does not match")
	}
}

func testFixture(name, expected string, t *testing.T) {
	in := getFixture(name, t)
	out := RemoveReplies(in)
	compare(expected, out, t)
}

func TestDoesNotModifyInputString(t *testing.T) {
	in := "The Quick Brown Fox Jumps Over The Lazy Dog"
	out := RemoveReplies(in)
	compare(in, out, t)
}

func TestEmailItalian(t *testing.T) {
	testFixture("email_7", CommonBody, t)
}

func TestEmailDutch(t *testing.T) {
	testFixture("email_8", CommonBody, t)
}

func TestEmailSignatureWithEqual(t *testing.T) {
	testFixture("email_9", CommonBody, t)
}

func TestEmailHotmail(t *testing.T) {
	testFixture("email_10", CommonBody, t)
}

func TestEmailWhitespaceBeforeHeader(t *testing.T) {
	testFixture("email_11", CommonBody, t)
}

func TestEmailWithSquareBrackets(t *testing.T) {
	testFixture("email_12", CommonBody, t)
}

func TestEmailDaIntoItalian(t *testing.T) {
	testFixture("email_13", CommonBody, t)
}

func TestEmailHeaderPolish(t *testing.T) {
	testFixture("email_14", CommonBody, t)
}

func TestEmailSentFromMy(t *testing.T) {
	testFixture("email_15", CommonBody, t)
}

func TestEmailHeaderPolishWithDniaAndNapisala(t *testing.T) {
	testFixture("email_16", CommonBody, t)
}

func TestEmailHeaderPolishWithDateInIso8601(t *testing.T) {
	testFixture("email_17", CommonBody, t)
}

func TestEmailOutlookEn(t *testing.T) {
	testFixture("email_18", CommonBody, t)
}

func TestEmailHeaderSimplifiedChinese(t *testing.T) {
	testFixture("email_22", CommonBody, t)
}

func TestEmailUkrainian(t *testing.T) {
	testFixture("email_23", CommonBody, t)
}

func TestEmailGmailNo(t *testing.T) {
	testFixture("email_norwegian_gmail", CommonBody, t)
}

func TestParseOutSentFromIPhone(t *testing.T) {
	testFixture("email_iphone", "Here is another email", t)
}

func TestParseOutSentFromBlackBerry(t *testing.T) {
	testFixture("email_blackberry", "Here is another email", t)
}

func TestParseOutSendFromMultiwordMobileDevice(t *testing.T) {
	testFixture("email_multi_word_sent_from_my_mobile_device", "Here is another email", t)
}

func TestDoNotParseOutSendFromInRegularSentence(t *testing.T) {
	testFixture(
		"email_sent_from_my_not_signature",
		"Here is another email\n\nSent from my desk, is much easier then my mobile phone.",
		t,
	)
}

func TestParseOutJustTopForOutlookReply(t *testing.T) {
	testFixture("email_2_1", "Outlook with a reply", t)
}

func TestRetainsBullets(t *testing.T) {
	testFixture(
		"email_bullets",
		"test 2 this should list second\n\nand have spaces\n\nand retain this formatting\n\n\n   - how about bullets\n   - and another",
		t,
	)
}

func TestUnquotedReply(t *testing.T) {
	testFixture("email_unquoted_reply", "This is my reply.", t)
}

func TestLiberkeysGSuiteSignature(t *testing.T) {
	testFixture("liberkeys_gsuite_signature", "Ceci est un email.", t)
}

func TestLiberkeysGSuiteSignature2(t *testing.T) {
	testFixture("liberkeys_gsuite_signature2", "Ceci est un email.", t)
}

func TestQuotedQuoteHeader(t *testing.T) {
	testFixture("quoted_quote_header", "Ceci est un email.", t)
}

func TestStrangeCaseWithoutReply(t *testing.T) {
	body := `Toto,

Le rendez vous est bien confirmé à 10h30. Vous retrouverez ma collaboratrice Tata Toto je vous joins son numéro si nécessaire 06 12 34 56 78.

Bien à vous,`
	testFixture("strange_case_without_reply", body, t)
}

func TestLiberkeysNewSignature(t *testing.T) {
	testFixture("liberkeys_new_signature", "Ceci est un email.", t)
}

func TestLiberkeysStrangeCase(t *testing.T) {
	body := `Bonjour Monsieur XXX,

Pour rappel, notre package à *2990**€* inclut:

*- Accompagnement complet, de l’estimation à l'acte authentique*
*- Partenariat pour les diagnostics ( -15€ avec le code LIBERDIAG
chez allodiagnostic.com <http://allodiagnostic.com/> ou depuis votre compte
Liberkeys)*
*- Shooting photo professionnel *
*- Diffusion immédiate des biens sur 40 portails avec budget marketing pour
être toujours en tête de liste*
*- Organisation et réalisation des visites*
*- Vérification de la fiabilité et de la solvabilité des acquéreurs*
*- Négociation des offres*
*- Accompagnement jusqu'à l'acte authentique.*

*=> Le tout via notre Web Application qui vous permet de suivre et
contrôler les ventes en toute transparence.*
*=> Seulement 2990**€ **au succès*

Blablabla`
	testFixture("liberkeys_strange_case", body, t)
}

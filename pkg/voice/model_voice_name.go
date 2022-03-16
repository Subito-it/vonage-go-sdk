/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.3.2
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice

// VoiceName The voice & language to use
type VoiceName string

// List of voice_name
const (
	ADITI      VoiceName = "Aditi"
	AGNIESZKA  VoiceName = "Agnieszka"
	ALVA       VoiceName = "Alva"
	AMY        VoiceName = "Amy"
	ASTRID     VoiceName = "Astrid"
	BIANCA     VoiceName = "Bianca"
	BRIAN      VoiceName = "Brian"
	CARLA      VoiceName = "Carla"
	CARMEN     VoiceName = "Carmen"
	CARMIT     VoiceName = "Carmit"
	CATARINA   VoiceName = "Catarina"
	CELINE     VoiceName = "Celine"
	CEM        VoiceName = "Cem"
	CHANTAL    VoiceName = "Chantal"
	CHIPMUNK   VoiceName = "Chipmunk"
	CONCHITA   VoiceName = "Conchita"
	CRISTIANO  VoiceName = "Cristiano"
	DAMAYANTI  VoiceName = "Damayanti"
	DORA       VoiceName = "Dora"
	EMMA       VoiceName = "Emma"
	EMPAR      VoiceName = "Empar"
	ENRIQUE    VoiceName = "Enrique"
	ERIC       VoiceName = "Eric"
	EWA        VoiceName = "Ewa"
	FELIPE     VoiceName = "Felipe"
	FILIZ      VoiceName = "Filiz"
	GERAINT    VoiceName = "Geraint"
	GIORGIO    VoiceName = "Giorgio"
	GWYNETH    VoiceName = "Gwyneth"
	HANS       VoiceName = "Hans"
	HENRIK     VoiceName = "Henrik"
	INES       VoiceName = "Ines"
	IOANA      VoiceName = "Ioana"
	IVETA      VoiceName = "Iveta"
	IVY        VoiceName = "Ivy"
	JACEK      VoiceName = "Jacek"
	JAN        VoiceName = "Jan"
	JENNIFER   VoiceName = "Jennifer"
	JOANA      VoiceName = "Joana"
	JOANNA     VoiceName = "Joanna"
	JOEY       VoiceName = "Joey"
	JORDI      VoiceName = "Jordi"
	JUSTIN     VoiceName = "Justin"
	KANYA      VoiceName = "Kanya"
	KARL       VoiceName = "Karl"
	KENDRA     VoiceName = "Kendra"
	KIMBERLY   VoiceName = "Kimberly"
	LAILA      VoiceName = "Laila"
	LAURA      VoiceName = "Laura"
	LEA        VoiceName = "Lea"
	LEKHA      VoiceName = "Lekha"
	LIV        VoiceName = "Liv"
	LOTTE      VoiceName = "Lotte"
	LUCIA      VoiceName = "Lucia"
	LUCIANA    VoiceName = "Luciana"
	MADS       VoiceName = "Mads"
	MAGED      VoiceName = "Maged"
	MAJA       VoiceName = "Maja"
	MARISKA    VoiceName = "Mariska"
	MARLENE    VoiceName = "Marlene"
	MATHIEU    VoiceName = "Mathieu"
	MATTHEW    VoiceName = "Matthew"
	MAXIM      VoiceName = "Maxim"
	MEI_JIA    VoiceName = "Mei-Jia"
	MELINA     VoiceName = "Melina"
	MIA        VoiceName = "Mia"
	MIGUEL     VoiceName = "Miguel"
	MIREN      VoiceName = "Miren"
	MIZUKI     VoiceName = "Mizuki"
	MONTSERRAT VoiceName = "Montserrat"
	NAJA       VoiceName = "Naja"
	NICOLE     VoiceName = "Nicole"
	NIKOS      VoiceName = "Nikos"
	NORA       VoiceName = "Nora"
	OSKAR      VoiceName = "Oskar"
	PENELOPE   VoiceName = "Penelope"
	RAVEENA    VoiceName = "Raveena"
	RICARDO    VoiceName = "Ricardo"
	RUBEN      VoiceName = "Ruben"
	RUSSELL    VoiceName = "Russell"
	SALLI      VoiceName = "Salli"
	SATU       VoiceName = "Satu"
	SEOYEON    VoiceName = "Seoyeon"
	SIN_JI     VoiceName = "Sin-Ji"
	SORA       VoiceName = "Sora"
	TAKUMI     VoiceName = "Takumi"
	TARIK      VoiceName = "Tarik"
	TATYANA    VoiceName = "Tatyana"
	TESSA      VoiceName = "Tessa"
	TIAN_TIAN  VoiceName = "Tian-Tian"
	VICKI      VoiceName = "Vicki"
	VITORIA    VoiceName = "Vitoria"
	YELDA      VoiceName = "Yelda"
	ZEINA      VoiceName = "Zeina"
	ZHIYU      VoiceName = "Zhiyu"
	ZUZANA     VoiceName = "Zuzana"
)

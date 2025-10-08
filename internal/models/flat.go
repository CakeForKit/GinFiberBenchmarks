package models

import (
	"encoding/json"
	"math/rand"
	"time"
)

func GenerateFlatStructJSON() ([]byte, error) {
	flatStruct := generateFlatStruct()
	return json.MarshalIndent(flatStruct, "", "  ")
}

type FlatStruct struct {
	Field1   string  `json:"field_1"`
	Field2   int     `json:"field_2"`
	Field3   float64 `json:"field_3"`
	Field4   bool    `json:"field_4"`
	Field5   string  `json:"field_5"`
	Field6   int     `json:"field_6"`
	Field7   float64 `json:"field_7"`
	Field8   bool    `json:"field_8"`
	Field9   string  `json:"field_9"`
	Field10  int     `json:"field_10"`
	Field11  float64 `json:"field_11"`
	Field12  bool    `json:"field_12"`
	Field13  string  `json:"field_13"`
	Field14  int     `json:"field_14"`
	Field15  float64 `json:"field_15"`
	Field16  bool    `json:"field_16"`
	Field17  string  `json:"field_17"`
	Field18  int     `json:"field_18"`
	Field19  float64 `json:"field_19"`
	Field20  bool    `json:"field_20"`
	Field21  string  `json:"field_21"`
	Field22  int     `json:"field_22"`
	Field23  float64 `json:"field_23"`
	Field24  bool    `json:"field_24"`
	Field25  string  `json:"field_25"`
	Field26  int     `json:"field_26"`
	Field27  float64 `json:"field_27"`
	Field28  bool    `json:"field_28"`
	Field29  string  `json:"field_29"`
	Field30  int     `json:"field_30"`
	Field31  float64 `json:"field_31"`
	Field32  bool    `json:"field_32"`
	Field33  string  `json:"field_33"`
	Field34  int     `json:"field_34"`
	Field35  float64 `json:"field_35"`
	Field36  bool    `json:"field_36"`
	Field37  string  `json:"field_37"`
	Field38  int     `json:"field_38"`
	Field39  float64 `json:"field_39"`
	Field40  bool    `json:"field_40"`
	Field41  string  `json:"field_41"`
	Field42  int     `json:"field_42"`
	Field43  float64 `json:"field_43"`
	Field44  bool    `json:"field_44"`
	Field45  string  `json:"field_45"`
	Field46  int     `json:"field_46"`
	Field47  float64 `json:"field_47"`
	Field48  bool    `json:"field_48"`
	Field49  string  `json:"field_49"`
	Field50  int     `json:"field_50"`
	Field51  float64 `json:"field_51"`
	Field52  bool    `json:"field_52"`
	Field53  string  `json:"field_53"`
	Field54  int     `json:"field_54"`
	Field55  float64 `json:"field_55"`
	Field56  bool    `json:"field_56"`
	Field57  string  `json:"field_57"`
	Field58  int     `json:"field_58"`
	Field59  float64 `json:"field_59"`
	Field60  bool    `json:"field_60"`
	Field61  string  `json:"field_61"`
	Field62  int     `json:"field_62"`
	Field63  float64 `json:"field_63"`
	Field64  bool    `json:"field_64"`
	Field65  string  `json:"field_65"`
	Field66  int     `json:"field_66"`
	Field67  float64 `json:"field_67"`
	Field68  bool    `json:"field_68"`
	Field69  string  `json:"field_69"`
	Field70  int     `json:"field_70"`
	Field71  float64 `json:"field_71"`
	Field72  bool    `json:"field_72"`
	Field73  string  `json:"field_73"`
	Field74  int     `json:"field_74"`
	Field75  float64 `json:"field_75"`
	Field76  bool    `json:"field_76"`
	Field77  string  `json:"field_77"`
	Field78  int     `json:"field_78"`
	Field79  float64 `json:"field_79"`
	Field80  bool    `json:"field_80"`
	Field81  string  `json:"field_81"`
	Field82  int     `json:"field_82"`
	Field83  float64 `json:"field_83"`
	Field84  bool    `json:"field_84"`
	Field85  string  `json:"field_85"`
	Field86  int     `json:"field_86"`
	Field87  float64 `json:"field_87"`
	Field88  bool    `json:"field_88"`
	Field89  string  `json:"field_89"`
	Field90  int     `json:"field_90"`
	Field91  float64 `json:"field_91"`
	Field92  bool    `json:"field_92"`
	Field93  string  `json:"field_93"`
	Field94  int     `json:"field_94"`
	Field95  float64 `json:"field_95"`
	Field96  bool    `json:"field_96"`
	Field97  string  `json:"field_97"`
	Field98  int     `json:"field_98"`
	Field99  float64 `json:"field_99"`
	Field100 bool    `json:"field_100"`
}

func generateFlatStruct() FlatStruct {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return FlatStruct{
		Field1:   randomString(8),
		Field2:   rand.Intn(1000),
		Field3:   rand.Float64() * 100,
		Field4:   rand.Intn(2) == 1,
		Field5:   randomString(10),
		Field6:   rand.Intn(500),
		Field7:   rand.Float64() * 50,
		Field8:   rand.Intn(2) == 1,
		Field9:   randomString(6),
		Field10:  rand.Intn(200),
		Field11:  rand.Float64() * 200,
		Field12:  rand.Intn(2) == 1,
		Field13:  randomString(12),
		Field14:  rand.Intn(800),
		Field15:  rand.Float64() * 150,
		Field16:  rand.Intn(2) == 1,
		Field17:  randomString(7),
		Field18:  rand.Intn(300),
		Field19:  rand.Float64() * 75,
		Field20:  rand.Intn(2) == 1,
		Field21:  randomString(9),
		Field22:  rand.Intn(400),
		Field23:  rand.Float64() * 125,
		Field24:  rand.Intn(2) == 1,
		Field25:  randomString(5),
		Field26:  rand.Intn(600),
		Field27:  rand.Float64() * 80,
		Field28:  rand.Intn(2) == 1,
		Field29:  randomString(11),
		Field30:  rand.Intn(700),
		Field31:  rand.Float64() * 90,
		Field32:  rand.Intn(2) == 1,
		Field33:  randomString(8),
		Field34:  rand.Intn(250),
		Field35:  rand.Float64() * 60,
		Field36:  rand.Intn(2) == 1,
		Field37:  randomString(10),
		Field38:  rand.Intn(350),
		Field39:  rand.Float64() * 110,
		Field40:  rand.Intn(2) == 1,
		Field41:  randomString(6),
		Field42:  rand.Intn(450),
		Field43:  rand.Float64() * 95,
		Field44:  rand.Intn(2) == 1,
		Field45:  randomString(12),
		Field46:  rand.Intn(550),
		Field47:  rand.Float64() * 130,
		Field48:  rand.Intn(2) == 1,
		Field49:  randomString(7),
		Field50:  rand.Intn(650),
		Field51:  rand.Float64() * 140,
		Field52:  rand.Intn(2) == 1,
		Field53:  randomString(9),
		Field54:  rand.Intn(750),
		Field55:  rand.Float64() * 160,
		Field56:  rand.Intn(2) == 1,
		Field57:  randomString(8),
		Field58:  rand.Intn(850),
		Field59:  rand.Float64() * 170,
		Field60:  rand.Intn(2) == 1,
		Field61:  randomString(10),
		Field62:  rand.Intn(950),
		Field63:  rand.Float64() * 180,
		Field64:  rand.Intn(2) == 1,
		Field65:  randomString(11),
		Field66:  rand.Intn(150),
		Field67:  rand.Float64() * 40,
		Field68:  rand.Intn(2) == 1,
		Field69:  randomString(6),
		Field70:  rand.Intn(220),
		Field71:  rand.Float64() * 70,
		Field72:  rand.Intn(2) == 1,
		Field73:  randomString(12),
		Field74:  rand.Intn(320),
		Field75:  rand.Float64() * 85,
		Field76:  rand.Intn(2) == 1,
		Field77:  randomString(7),
		Field78:  rand.Intn(420),
		Field79:  rand.Float64() * 105,
		Field80:  rand.Intn(2) == 1,
		Field81:  randomString(9),
		Field82:  rand.Intn(520),
		Field83:  rand.Float64() * 115,
		Field84:  rand.Intn(2) == 1,
		Field85:  randomString(8),
		Field86:  rand.Intn(620),
		Field87:  rand.Float64() * 135,
		Field88:  rand.Intn(2) == 1,
		Field89:  randomString(10),
		Field90:  rand.Intn(720),
		Field91:  rand.Float64() * 145,
		Field92:  rand.Intn(2) == 1,
		Field93:  randomString(11),
		Field94:  rand.Intn(820),
		Field95:  rand.Float64() * 155,
		Field96:  rand.Intn(2) == 1,
		Field97:  randomString(6),
		Field98:  rand.Intn(920),
		Field99:  rand.Float64() * 165,
		Field100: rand.Intn(2) == 1,
	}
}

// Вспомогательная функция для генерации случайных строк
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

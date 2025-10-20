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
	Field101 string  `json:"field_101"`
	Field102 int     `json:"field_102"`
	Field103 float64 `json:"field_103"`
	Field104 bool    `json:"field_104"`
	Field105 string  `json:"field_105"`
	Field106 int     `json:"field_106"`
	Field107 float64 `json:"field_107"`
	Field108 bool    `json:"field_108"`
	Field109 string  `json:"field_109"`
	Field110 int     `json:"field_110"`
	Field111 float64 `json:"field_111"`
	Field112 bool    `json:"field_112"`
	Field113 string  `json:"field_113"`
	Field114 int     `json:"field_114"`
	Field115 float64 `json:"field_115"`
	Field116 bool    `json:"field_116"`
	Field117 string  `json:"field_117"`
	Field118 int     `json:"field_118"`
	Field119 float64 `json:"field_119"`
	Field120 bool    `json:"field_120"`
	Field121 string  `json:"field_121"`
	Field122 int     `json:"field_122"`
	Field123 float64 `json:"field_123"`
	Field124 bool    `json:"field_124"`
	Field125 string  `json:"field_125"`
	Field126 int     `json:"field_126"`
	Field127 float64 `json:"field_127"`
	Field128 bool    `json:"field_128"`
	Field129 string  `json:"field_129"`
	Field130 int     `json:"field_130"`
	Field131 float64 `json:"field_131"`
	Field132 bool    `json:"field_132"`
	Field133 string  `json:"field_133"`
	Field134 int     `json:"field_134"`
	Field135 float64 `json:"field_135"`
	Field136 bool    `json:"field_136"`
	Field137 string  `json:"field_137"`
	Field138 int     `json:"field_138"`
	Field139 float64 `json:"field_139"`
	Field140 bool    `json:"field_140"`
	Field141 string  `json:"field_141"`
	Field142 int     `json:"field_142"`
	Field143 float64 `json:"field_143"`
	Field144 bool    `json:"field_144"`
	Field145 string  `json:"field_145"`
	Field146 int     `json:"field_146"`
	Field147 float64 `json:"field_147"`
	Field148 bool    `json:"field_148"`
	Field149 string  `json:"field_149"`
	Field150 int     `json:"field_150"`
	Field151 float64 `json:"field_151"`
	Field152 bool    `json:"field_152"`
	Field153 string  `json:"field_153"`
	Field154 int     `json:"field_154"`
	Field155 float64 `json:"field_155"`
	Field156 bool    `json:"field_156"`
	Field157 string  `json:"field_157"`
	Field158 int     `json:"field_158"`
	Field159 float64 `json:"field_159"`
	Field160 bool    `json:"field_160"`
	Field161 string  `json:"field_161"`
	Field162 int     `json:"field_162"`
	Field163 float64 `json:"field_163"`
	Field164 bool    `json:"field_164"`
	Field165 string  `json:"field_165"`
	Field166 int     `json:"field_166"`
	Field167 float64 `json:"field_167"`
	Field168 bool    `json:"field_168"`
	Field169 string  `json:"field_169"`
	Field170 int     `json:"field_170"`
	Field171 float64 `json:"field_171"`
	Field172 bool    `json:"field_172"`
	Field173 string  `json:"field_173"`
	Field174 int     `json:"field_174"`
	Field175 float64 `json:"field_175"`
	Field176 bool    `json:"field_176"`
	Field177 string  `json:"field_177"`
	Field178 int     `json:"field_178"`
	Field179 float64 `json:"field_179"`
	Field180 bool    `json:"field_180"`
	Field181 string  `json:"field_181"`
	Field182 int     `json:"field_182"`
	Field183 float64 `json:"field_183"`
	Field184 bool    `json:"field_184"`
	Field185 string  `json:"field_185"`
	Field186 int     `json:"field_186"`
	Field187 float64 `json:"field_187"`
	Field188 bool    `json:"field_188"`
	Field189 string  `json:"field_189"`
	Field190 int     `json:"field_190"`
	Field191 float64 `json:"field_191"`
	Field192 bool    `json:"field_192"`
	Field193 string  `json:"field_193"`
	Field194 int     `json:"field_194"`
	Field195 float64 `json:"field_195"`
	Field196 bool    `json:"field_196"`
	Field197 string  `json:"field_197"`
	Field198 int     `json:"field_198"`
	Field199 float64 `json:"field_199"`
	Field200 bool    `json:"field_200"`
	Field201 string  `json:"field_201"`
	Field202 int     `json:"field_202"`
	Field203 float64 `json:"field_203"`
	Field204 bool    `json:"field_204"`
	Field205 string  `json:"field_205"`
	Field206 int     `json:"field_206"`
	Field207 float64 `json:"field_207"`
	Field208 bool    `json:"field_208"`
	Field209 string  `json:"field_209"`
	Field210 int     `json:"field_210"`
	Field211 float64 `json:"field_211"`
	Field212 bool    `json:"field_212"`
	Field213 string  `json:"field_213"`
	Field214 int     `json:"field_214"`
	Field215 float64 `json:"field_215"`
	Field216 bool    `json:"field_216"`
	Field217 string  `json:"field_217"`
	Field218 int     `json:"field_218"`
	Field219 float64 `json:"field_219"`
	Field220 bool    `json:"field_220"`
	Field221 string  `json:"field_221"`
	Field222 int     `json:"field_222"`
	Field223 float64 `json:"field_223"`
	Field224 bool    `json:"field_224"`
	Field225 string  `json:"field_225"`
	Field226 int     `json:"field_226"`
	Field227 float64 `json:"field_227"`
	Field228 bool    `json:"field_228"`
	Field229 string  `json:"field_229"`
	Field230 int     `json:"field_230"`
	Field231 float64 `json:"field_231"`
	Field232 bool    `json:"field_232"`
	Field233 string  `json:"field_233"`
	Field234 int     `json:"field_234"`
	Field235 float64 `json:"field_235"`
	Field236 bool    `json:"field_236"`
	Field237 string  `json:"field_237"`
	Field238 int     `json:"field_238"`
	Field239 float64 `json:"field_239"`
	Field240 bool    `json:"field_240"`
	Field241 string  `json:"field_241"`
	Field242 int     `json:"field_242"`
	Field243 float64 `json:"field_243"`
	Field244 bool    `json:"field_244"`
	Field245 string  `json:"field_245"`
	Field246 int     `json:"field_246"`
	Field247 float64 `json:"field_247"`
	Field248 bool    `json:"field_248"`
	Field249 string  `json:"field_249"`
	Field250 int     `json:"field_250"`
	Field251 float64 `json:"field_251"`
	Field252 bool    `json:"field_252"`
	Field253 string  `json:"field_253"`
	Field254 int     `json:"field_254"`
	Field255 float64 `json:"field_255"`
	Field256 bool    `json:"field_256"`
	Field257 string  `json:"field_257"`
	Field258 int     `json:"field_258"`
	Field259 float64 `json:"field_259"`
	Field260 bool    `json:"field_260"`
	Field261 string  `json:"field_261"`
	Field262 int     `json:"field_262"`
	Field263 float64 `json:"field_263"`
	Field264 bool    `json:"field_264"`
	Field265 string  `json:"field_265"`
	Field266 int     `json:"field_266"`
	Field267 float64 `json:"field_267"`
	Field268 bool    `json:"field_268"`
	Field269 string  `json:"field_269"`
	Field270 int     `json:"field_270"`
	Field271 float64 `json:"field_271"`
	Field272 bool    `json:"field_272"`
	Field273 string  `json:"field_273"`
	Field274 int     `json:"field_274"`
	Field275 float64 `json:"field_275"`
	Field276 bool    `json:"field_276"`
	Field277 string  `json:"field_277"`
	Field278 int     `json:"field_278"`
	Field279 float64 `json:"field_279"`
	Field280 bool    `json:"field_280"`
	Field281 string  `json:"field_281"`
	Field282 int     `json:"field_282"`
	Field283 float64 `json:"field_283"`
	Field284 bool    `json:"field_284"`
	Field285 string  `json:"field_285"`
	Field286 int     `json:"field_286"`
	Field287 float64 `json:"field_287"`
	Field288 bool    `json:"field_288"`
	Field289 string  `json:"field_289"`
	Field290 int     `json:"field_290"`
	Field291 float64 `json:"field_291"`
	Field292 bool    `json:"field_292"`
	Field293 string  `json:"field_293"`
	Field294 int     `json:"field_294"`
	Field295 float64 `json:"field_295"`
	Field296 bool    `json:"field_296"`
	Field297 string  `json:"field_297"`
	Field298 int     `json:"field_298"`
	Field299 float64 `json:"field_299"`
	Field300 bool    `json:"field_300"`
	Field301 string  `json:"field_301"`
	Field302 int     `json:"field_302"`
	Field303 float64 `json:"field_303"`
	Field304 bool    `json:"field_304"`
	Field305 string  `json:"field_305"`
	Field306 int     `json:"field_306"`
	Field307 float64 `json:"field_307"`
	Field308 bool    `json:"field_308"`
	Field309 string  `json:"field_309"`
	Field310 int     `json:"field_310"`
	Field311 float64 `json:"field_311"`
	Field312 bool    `json:"field_312"`
	Field313 string  `json:"field_313"`
	Field314 int     `json:"field_314"`
	Field315 float64 `json:"field_315"`
	Field316 bool    `json:"field_316"`
	Field317 string  `json:"field_317"`
	Field318 int     `json:"field_318"`
	Field319 float64 `json:"field_319"`
	Field320 bool    `json:"field_320"`
	Field321 string  `json:"field_321"`
	Field322 int     `json:"field_322"`
	Field323 float64 `json:"field_323"`
	Field324 bool    `json:"field_324"`
	Field325 string  `json:"field_325"`
	Field326 int     `json:"field_326"`
	Field327 float64 `json:"field_327"`
	Field328 bool    `json:"field_328"`
	Field329 string  `json:"field_329"`
	Field330 int     `json:"field_330"`
	Field331 float64 `json:"field_331"`
	Field332 bool    `json:"field_332"`
	Field333 string  `json:"field_333"`
	Field334 int     `json:"field_334"`
	Field335 float64 `json:"field_335"`
	Field336 bool    `json:"field_336"`
	Field337 string  `json:"field_337"`
	Field338 int     `json:"field_338"`
	Field339 float64 `json:"field_339"`
	Field340 bool    `json:"field_340"`
	Field341 string  `json:"field_341"`
	Field342 int     `json:"field_342"`
	Field343 float64 `json:"field_343"`
	Field344 bool    `json:"field_344"`
	Field345 string  `json:"field_345"`
	Field346 int     `json:"field_346"`
	Field347 float64 `json:"field_347"`
	Field348 bool    `json:"field_348"`
	Field349 string  `json:"field_349"`
	Field350 int     `json:"field_350"`
	Field351 float64 `json:"field_351"`
	Field352 bool    `json:"field_352"`
	Field353 string  `json:"field_353"`
	Field354 int     `json:"field_354"`
	Field355 float64 `json:"field_355"`
	Field356 bool    `json:"field_356"`
	Field357 string  `json:"field_357"`
	Field358 int     `json:"field_358"`
	Field359 float64 `json:"field_359"`
	Field360 bool    `json:"field_360"`
	Field361 string  `json:"field_361"`
	Field362 int     `json:"field_362"`
	Field363 float64 `json:"field_363"`
	Field364 bool    `json:"field_364"`
	Field365 string  `json:"field_365"`
	Field366 int     `json:"field_366"`
	Field367 float64 `json:"field_367"`
	Field368 bool    `json:"field_368"`
	Field369 string  `json:"field_369"`
	Field370 int     `json:"field_370"`
	Field371 float64 `json:"field_371"`
	Field372 bool    `json:"field_372"`
	Field373 string  `json:"field_373"`
	Field374 int     `json:"field_374"`
	Field375 float64 `json:"field_375"`
	Field376 bool    `json:"field_376"`
	Field377 string  `json:"field_377"`
	Field378 int     `json:"field_378"`
	Field379 float64 `json:"field_379"`
	Field380 bool    `json:"field_380"`
	Field381 string  `json:"field_381"`
	Field382 int     `json:"field_382"`
	Field383 float64 `json:"field_383"`
	Field384 bool    `json:"field_384"`
	Field385 string  `json:"field_385"`
	Field386 int     `json:"field_386"`
	Field387 float64 `json:"field_387"`
	Field388 bool    `json:"field_388"`
	Field389 string  `json:"field_389"`
	Field390 int     `json:"field_390"`
	Field391 float64 `json:"field_391"`
	Field392 bool    `json:"field_392"`
	Field393 string  `json:"field_393"`
	Field394 int     `json:"field_394"`
	Field395 float64 `json:"field_395"`
	Field396 bool    `json:"field_396"`
	Field397 string  `json:"field_397"`
	Field398 int     `json:"field_398"`
	Field399 float64 `json:"field_399"`
	Field400 bool    `json:"field_400"`
	Field401 string  `json:"field_401"`
	Field402 int     `json:"field_402"`
	Field403 float64 `json:"field_403"`
	Field404 bool    `json:"field_404"`
	Field405 string  `json:"field_405"`
	Field406 int     `json:"field_406"`
	Field407 float64 `json:"field_407"`
	Field408 bool    `json:"field_408"`
	Field409 string  `json:"field_409"`
	Field410 int     `json:"field_410"`
	Field411 float64 `json:"field_411"`
	Field412 bool    `json:"field_412"`
	Field413 string  `json:"field_413"`
	Field414 int     `json:"field_414"`
	Field415 float64 `json:"field_415"`
	Field416 bool    `json:"field_416"`
	Field417 string  `json:"field_417"`
	Field418 int     `json:"field_418"`
	Field419 float64 `json:"field_419"`
	Field420 bool    `json:"field_420"`
	Field421 string  `json:"field_421"`
	Field422 int     `json:"field_422"`
	Field423 float64 `json:"field_423"`
	Field424 bool    `json:"field_424"`
	Field425 string  `json:"field_425"`
	Field426 int     `json:"field_426"`
	Field427 float64 `json:"field_427"`
	Field428 bool    `json:"field_428"`
	Field429 string  `json:"field_429"`
	Field430 int     `json:"field_430"`
	Field431 float64 `json:"field_431"`
	Field432 bool    `json:"field_432"`
	Field433 string  `json:"field_433"`
	Field434 int     `json:"field_434"`
	Field435 float64 `json:"field_435"`
	Field436 bool    `json:"field_436"`
	Field437 string  `json:"field_437"`
	Field438 int     `json:"field_438"`
	Field439 float64 `json:"field_439"`
	Field440 bool    `json:"field_440"`
	Field441 string  `json:"field_441"`
	Field442 int     `json:"field_442"`
	Field443 float64 `json:"field_443"`
	Field444 bool    `json:"field_444"`
	Field445 string  `json:"field_445"`
	Field446 int     `json:"field_446"`
	Field447 float64 `json:"field_447"`
	Field448 bool    `json:"field_448"`
	Field449 string  `json:"field_449"`
	Field450 int     `json:"field_450"`
	Field451 float64 `json:"field_451"`
	Field452 bool    `json:"field_452"`
	Field453 string  `json:"field_453"`
	Field454 int     `json:"field_454"`
	Field455 float64 `json:"field_455"`
	Field456 bool    `json:"field_456"`
	Field457 string  `json:"field_457"`
	Field458 int     `json:"field_458"`
	Field459 float64 `json:"field_459"`
	Field460 bool    `json:"field_460"`
	Field461 string  `json:"field_461"`
	Field462 int     `json:"field_462"`
	Field463 float64 `json:"field_463"`
	Field464 bool    `json:"field_464"`
	Field465 string  `json:"field_465"`
	Field466 int     `json:"field_466"`
	Field467 float64 `json:"field_467"`
	Field468 bool    `json:"field_468"`
	Field469 string  `json:"field_469"`
	Field470 int     `json:"field_470"`
	Field471 float64 `json:"field_471"`
	Field472 bool    `json:"field_472"`
	Field473 string  `json:"field_473"`
	Field474 int     `json:"field_474"`
	Field475 float64 `json:"field_475"`
	Field476 bool    `json:"field_476"`
	Field477 string  `json:"field_477"`
	Field478 int     `json:"field_478"`
	Field479 float64 `json:"field_479"`
	Field480 bool    `json:"field_480"`
	Field481 string  `json:"field_481"`
	Field482 int     `json:"field_482"`
	Field483 float64 `json:"field_483"`
	Field484 bool    `json:"field_484"`
	Field485 string  `json:"field_485"`
	Field486 int     `json:"field_486"`
	Field487 float64 `json:"field_487"`
	Field488 bool    `json:"field_488"`
	Field489 string  `json:"field_489"`
	Field490 int     `json:"field_490"`
	Field491 float64 `json:"field_491"`
	Field492 bool    `json:"field_492"`
	Field493 string  `json:"field_493"`
	Field494 int     `json:"field_494"`
	Field495 float64 `json:"field_495"`
	Field496 bool    `json:"field_496"`
	Field497 string  `json:"field_497"`
	Field498 int     `json:"field_498"`
	Field499 float64 `json:"field_499"`
	Field500 bool    `json:"field_500"`
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
		Field101: randomString(8),
		Field102: rand.Intn(1000),
		Field103: rand.Float64() * 100,
		Field104: rand.Intn(2) == 1,
		Field105: randomString(10),
		Field106: rand.Intn(500),
		Field107: rand.Float64() * 50,
		Field108: rand.Intn(2) == 1,
		Field109: randomString(6),
		Field110: rand.Intn(200),
		Field111: rand.Float64() * 200,
		Field112: rand.Intn(2) == 1,
		Field113: randomString(12),
		Field114: rand.Intn(800),
		Field115: rand.Float64() * 150,
		Field116: rand.Intn(2) == 1,
		Field117: randomString(7),
		Field118: rand.Intn(300),
		Field119: rand.Float64() * 75,
		Field120: rand.Intn(2) == 1,
		Field121: randomString(9),
		Field122: rand.Intn(400),
		Field123: rand.Float64() * 125,
		Field124: rand.Intn(2) == 1,
		Field125: randomString(5),
		Field126: rand.Intn(600),
		Field127: rand.Float64() * 80,
		Field128: rand.Intn(2) == 1,
		Field129: randomString(11),
		Field130: rand.Intn(700),
		Field131: rand.Float64() * 90,
		Field132: rand.Intn(2) == 1,
		Field133: randomString(8),
		Field134: rand.Intn(250),
		Field135: rand.Float64() * 60,
		Field136: rand.Intn(2) == 1,
		Field137: randomString(10),
		Field138: rand.Intn(350),
		Field139: rand.Float64() * 110,
		Field140: rand.Intn(2) == 1,
		Field141: randomString(6),
		Field142: rand.Intn(450),
		Field143: rand.Float64() * 95,
		Field144: rand.Intn(2) == 1,
		Field145: randomString(12),
		Field146: rand.Intn(550),
		Field147: rand.Float64() * 130,
		Field148: rand.Intn(2) == 1,
		Field149: randomString(7),
		Field150: rand.Intn(650),
		Field151: rand.Float64() * 140,
		Field152: rand.Intn(2) == 1,
		Field153: randomString(9),
		Field154: rand.Intn(750),
		Field155: rand.Float64() * 160,
		Field156: rand.Intn(2) == 1,
		Field157: randomString(8),
		Field158: rand.Intn(850),
		Field159: rand.Float64() * 170,
		Field160: rand.Intn(2) == 1,
		Field161: randomString(10),
		Field162: rand.Intn(950),
		Field163: rand.Float64() * 180,
		Field164: rand.Intn(2) == 1,
		Field165: randomString(11),
		Field166: rand.Intn(150),
		Field167: rand.Float64() * 40,
		Field168: rand.Intn(2) == 1,
		Field169: randomString(6),
		Field170: rand.Intn(220),
		Field171: rand.Float64() * 70,
		Field172: rand.Intn(2) == 1,
		Field173: randomString(12),
		Field174: rand.Intn(320),
		Field175: rand.Float64() * 85,
		Field176: rand.Intn(2) == 1,
		Field177: randomString(7),
		Field178: rand.Intn(420),
		Field179: rand.Float64() * 105,
		Field180: rand.Intn(2) == 1,
		Field181: randomString(9),
		Field182: rand.Intn(520),
		Field183: rand.Float64() * 115,
		Field184: rand.Intn(2) == 1,
		Field185: randomString(8),
		Field186: rand.Intn(620),
		Field187: rand.Float64() * 135,
		Field188: rand.Intn(2) == 1,
		Field189: randomString(10),
		Field190: rand.Intn(720),
		Field191: rand.Float64() * 145,
		Field192: rand.Intn(2) == 1,
		Field193: randomString(11),
		Field194: rand.Intn(820),
		Field195: rand.Float64() * 155,
		Field196: rand.Intn(2) == 1,
		Field197: randomString(6),
		Field198: rand.Intn(920),
		Field199: rand.Float64() * 165,
		Field200: rand.Intn(2) == 1,
		Field201: randomString(8),
		Field202: rand.Intn(1000),
		Field203: rand.Float64() * 100,
		Field204: rand.Intn(2) == 1,
		Field205: randomString(10),
		Field206: rand.Intn(500),
		Field207: rand.Float64() * 50,
		Field208: rand.Intn(2) == 1,
		Field209: randomString(6),
		Field210: rand.Intn(200),
		Field211: rand.Float64() * 200,
		Field212: rand.Intn(2) == 1,
		Field213: randomString(12),
		Field214: rand.Intn(800),
		Field215: rand.Float64() * 150,
		Field216: rand.Intn(2) == 1,
		Field217: randomString(7),
		Field218: rand.Intn(300),
		Field219: rand.Float64() * 75,
		Field220: rand.Intn(2) == 1,
		Field221: randomString(9),
		Field222: rand.Intn(400),
		Field223: rand.Float64() * 125,
		Field224: rand.Intn(2) == 1,
		Field225: randomString(5),
		Field226: rand.Intn(600),
		Field227: rand.Float64() * 80,
		Field228: rand.Intn(2) == 1,
		Field229: randomString(11),
		Field230: rand.Intn(700),
		Field231: rand.Float64() * 90,
		Field232: rand.Intn(2) == 1,
		Field233: randomString(8),
		Field234: rand.Intn(250),
		Field235: rand.Float64() * 60,
		Field236: rand.Intn(2) == 1,
		Field237: randomString(10),
		Field238: rand.Intn(350),
		Field239: rand.Float64() * 110,
		Field240: rand.Intn(2) == 1,
		Field241: randomString(6),
		Field242: rand.Intn(450),
		Field243: rand.Float64() * 95,
		Field244: rand.Intn(2) == 1,
		Field245: randomString(12),
		Field246: rand.Intn(550),
		Field247: rand.Float64() * 130,
		Field248: rand.Intn(2) == 1,
		Field249: randomString(7),
		Field250: rand.Intn(650),
		Field251: rand.Float64() * 140,
		Field252: rand.Intn(2) == 1,
		Field253: randomString(9),
		Field254: rand.Intn(750),
		Field255: rand.Float64() * 160,
		Field256: rand.Intn(2) == 1,
		Field257: randomString(8),
		Field258: rand.Intn(850),
		Field259: rand.Float64() * 170,
		Field260: rand.Intn(2) == 1,
		Field261: randomString(10),
		Field262: rand.Intn(950),
		Field263: rand.Float64() * 180,
		Field264: rand.Intn(2) == 1,
		Field265: randomString(11),
		Field266: rand.Intn(150),
		Field267: rand.Float64() * 40,
		Field268: rand.Intn(2) == 1,
		Field269: randomString(6),
		Field270: rand.Intn(220),
		Field271: rand.Float64() * 70,
		Field272: rand.Intn(2) == 1,
		Field273: randomString(12),
		Field274: rand.Intn(320),
		Field275: rand.Float64() * 85,
		Field276: rand.Intn(2) == 1,
		Field277: randomString(7),
		Field278: rand.Intn(420),
		Field279: rand.Float64() * 105,
		Field280: rand.Intn(2) == 1,
		Field281: randomString(9),
		Field282: rand.Intn(520),
		Field283: rand.Float64() * 115,
		Field284: rand.Intn(2) == 1,
		Field285: randomString(8),
		Field286: rand.Intn(620),
		Field287: rand.Float64() * 135,
		Field288: rand.Intn(2) == 1,
		Field289: randomString(10),
		Field290: rand.Intn(720),
		Field291: rand.Float64() * 145,
		Field292: rand.Intn(2) == 1,
		Field293: randomString(11),
		Field294: rand.Intn(820),
		Field295: rand.Float64() * 155,
		Field296: rand.Intn(2) == 1,
		Field297: randomString(6),
		Field298: rand.Intn(920),
		Field299: rand.Float64() * 165,
		Field300: rand.Intn(2) == 1,
		Field301: randomString(8),
		Field302: rand.Intn(1000),
		Field303: rand.Float64() * 100,
		Field304: rand.Intn(2) == 1,
		Field305: randomString(10),
		Field306: rand.Intn(500),
		Field307: rand.Float64() * 50,
		Field308: rand.Intn(2) == 1,
		Field309: randomString(6),
		Field310: rand.Intn(200),
		Field311: rand.Float64() * 200,
		Field312: rand.Intn(2) == 1,
		Field313: randomString(12),
		Field314: rand.Intn(800),
		Field315: rand.Float64() * 150,
		Field316: rand.Intn(2) == 1,
		Field317: randomString(7),
		Field318: rand.Intn(300),
		Field319: rand.Float64() * 75,
		Field320: rand.Intn(2) == 1,
		Field321: randomString(9),
		Field322: rand.Intn(400),
		Field323: rand.Float64() * 125,
		Field324: rand.Intn(2) == 1,
		Field325: randomString(5),
		Field326: rand.Intn(600),
		Field327: rand.Float64() * 80,
		Field328: rand.Intn(2) == 1,
		Field329: randomString(11),
		Field330: rand.Intn(700),
		Field331: rand.Float64() * 90,
		Field332: rand.Intn(2) == 1,
		Field333: randomString(8),
		Field334: rand.Intn(250),
		Field335: rand.Float64() * 60,
		Field336: rand.Intn(2) == 1,
		Field337: randomString(10),
		Field338: rand.Intn(350),
		Field339: rand.Float64() * 110,
		Field340: rand.Intn(2) == 1,
		Field341: randomString(6),
		Field342: rand.Intn(450),
		Field343: rand.Float64() * 95,
		Field344: rand.Intn(2) == 1,
		Field345: randomString(12),
		Field346: rand.Intn(550),
		Field347: rand.Float64() * 130,
		Field348: rand.Intn(2) == 1,
		Field349: randomString(7),
		Field350: rand.Intn(650),
		Field351: rand.Float64() * 140,
		Field352: rand.Intn(2) == 1,
		Field353: randomString(9),
		Field354: rand.Intn(750),
		Field355: rand.Float64() * 160,
		Field356: rand.Intn(2) == 1,
		Field357: randomString(8),
		Field358: rand.Intn(850),
		Field359: rand.Float64() * 170,
		Field360: rand.Intn(2) == 1,
		Field361: randomString(10),
		Field362: rand.Intn(950),
		Field363: rand.Float64() * 180,
		Field364: rand.Intn(2) == 1,
		Field365: randomString(11),
		Field366: rand.Intn(150),
		Field367: rand.Float64() * 40,
		Field368: rand.Intn(2) == 1,
		Field369: randomString(6),
		Field370: rand.Intn(220),
		Field371: rand.Float64() * 70,
		Field372: rand.Intn(2) == 1,
		Field373: randomString(12),
		Field374: rand.Intn(320),
		Field375: rand.Float64() * 85,
		Field376: rand.Intn(2) == 1,
		Field377: randomString(7),
		Field378: rand.Intn(420),
		Field379: rand.Float64() * 105,
		Field380: rand.Intn(2) == 1,
		Field381: randomString(9),
		Field382: rand.Intn(520),
		Field383: rand.Float64() * 115,
		Field384: rand.Intn(2) == 1,
		Field385: randomString(8),
		Field386: rand.Intn(620),
		Field387: rand.Float64() * 135,
		Field388: rand.Intn(2) == 1,
		Field389: randomString(10),
		Field390: rand.Intn(720),
		Field391: rand.Float64() * 145,
		Field392: rand.Intn(2) == 1,
		Field393: randomString(11),
		Field394: rand.Intn(820),
		Field395: rand.Float64() * 155,
		Field396: rand.Intn(2) == 1,
		Field397: randomString(6),
		Field398: rand.Intn(920),
		Field399: rand.Float64() * 165,
		Field400: rand.Intn(2) == 1,
		Field401: randomString(8),
		Field402: rand.Intn(1000),
		Field403: rand.Float64() * 100,
		Field404: rand.Intn(2) == 1,
		Field405: randomString(10),
		Field406: rand.Intn(500),
		Field407: rand.Float64() * 50,
		Field408: rand.Intn(2) == 1,
		Field409: randomString(6),
		Field410: rand.Intn(200),
		Field411: rand.Float64() * 200,
		Field412: rand.Intn(2) == 1,
		Field413: randomString(12),
		Field414: rand.Intn(800),
		Field415: rand.Float64() * 150,
		Field416: rand.Intn(2) == 1,
		Field417: randomString(7),
		Field418: rand.Intn(300),
		Field419: rand.Float64() * 75,
		Field420: rand.Intn(2) == 1,
		Field421: randomString(9),
		Field422: rand.Intn(400),
		Field423: rand.Float64() * 125,
		Field424: rand.Intn(2) == 1,
		Field425: randomString(5),
		Field426: rand.Intn(600),
		Field427: rand.Float64() * 80,
		Field428: rand.Intn(2) == 1,
		Field429: randomString(11),
		Field430: rand.Intn(700),
		Field431: rand.Float64() * 90,
		Field432: rand.Intn(2) == 1,
		Field433: randomString(8),
		Field434: rand.Intn(250),
		Field435: rand.Float64() * 60,
		Field436: rand.Intn(2) == 1,
		Field437: randomString(10),
		Field438: rand.Intn(350),
		Field439: rand.Float64() * 110,
		Field440: rand.Intn(2) == 1,
		Field441: randomString(6),
		Field442: rand.Intn(450),
		Field443: rand.Float64() * 95,
		Field444: rand.Intn(2) == 1,
		Field445: randomString(12),
		Field446: rand.Intn(550),
		Field447: rand.Float64() * 130,
		Field448: rand.Intn(2) == 1,
		Field449: randomString(7),
		Field450: rand.Intn(650),
		Field451: rand.Float64() * 140,
		Field452: rand.Intn(2) == 1,
		Field453: randomString(9),
		Field454: rand.Intn(750),
		Field455: rand.Float64() * 160,
		Field456: rand.Intn(2) == 1,
		Field457: randomString(8),
		Field458: rand.Intn(850),
		Field459: rand.Float64() * 170,
		Field460: rand.Intn(2) == 1,
		Field461: randomString(10),
		Field462: rand.Intn(950),
		Field463: rand.Float64() * 180,
		Field464: rand.Intn(2) == 1,
		Field465: randomString(11),
		Field466: rand.Intn(150),
		Field467: rand.Float64() * 40,
		Field468: rand.Intn(2) == 1,
		Field469: randomString(6),
		Field470: rand.Intn(220),
		Field471: rand.Float64() * 70,
		Field472: rand.Intn(2) == 1,
		Field473: randomString(12),
		Field474: rand.Intn(320),
		Field475: rand.Float64() * 85,
		Field476: rand.Intn(2) == 1,
		Field477: randomString(7),
		Field478: rand.Intn(420),
		Field479: rand.Float64() * 105,
		Field480: rand.Intn(2) == 1,
		Field481: randomString(9),
		Field482: rand.Intn(520),
		Field483: rand.Float64() * 115,
		Field484: rand.Intn(2) == 1,
		Field485: randomString(8),
		Field486: rand.Intn(620),
		Field487: rand.Float64() * 135,
		Field488: rand.Intn(2) == 1,
		Field489: randomString(10),
		Field490: rand.Intn(720),
		Field491: rand.Float64() * 145,
		Field492: rand.Intn(2) == 1,
		Field493: randomString(11),
		Field494: rand.Intn(820),
		Field495: rand.Float64() * 155,
		Field496: rand.Intn(2) == 1,
		Field497: randomString(6),
		Field498: rand.Intn(920),
		Field499: rand.Float64() * 165,
		Field500: rand.Intn(2) == 1,
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

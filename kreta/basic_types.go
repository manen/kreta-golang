package kreta

import "time"

type GradeType struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type GradeMode struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type SubjectCategory struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type Subject struct {
	ID       string          `json:"Uid"`
	Category SubjectCategory `json:"Kategoria"`
	Title    string          `json:"Nev"`
}

type ClassGroup struct {
	ID string `json:"Uid"`
}

type Grade struct {
	ID               string     `json:"Uid"`
	TeacherName      string     `json:"ErtekeloTanarNeve"`
	GradeType        GradeType  `json:"ErtekFajta"`
	Kind             string     `json:"Jelleg"`
	DateRecorded     time.Time  `json:"KeszitesDatuma"`
	DateSeen         time.Time  `json:"LattamozasDatuma"`
	Mode             GradeMode  `json:"Mod"`
	Date             time.Time  `json:"RogzitesDatuma"`
	Weight           int        `json:"SulySzazalekErteke"`
	Value            int        `json:"SzamErtek"`
	ValueString      string     `json:"SzovegesErtek"`
	ValueStringShort string     `json:"SzovegesErtekelesRovidNev"`
	Subject          Subject    `json:"Tantargy"`
	Topic            string     `json:"Tema"`
	Type             GradeType  `json:"Tipus"`
	ClassGroup       ClassGroup `json:"OsztalyCsoport"`
}

type AbsenceCertType struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type AbsenceMode struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type AbsenceLesson struct {
	From  time.Time `json:"KezdoDatum"`
	To    time.Time `json:"VegDatum"`
	Hours int       `json:"Oraszam"`
}

type AbsenceType struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type Absence struct {
	CertStatus string          `json:"IgazolasAllapota"`
	CertType   AbsenceCertType `json:"IgazolasTipusa"`
	// DelayInMinutes type assumed from context
	DelayInMinutes int           `json:"KesesPercben"`
	DateRecorded   time.Time     `json:"KeszitesDatuma"`
	Mode           AbsenceMode   `json:"Mod"`
	Date           time.Time     `json:"Datum"`
	Lesson         AbsenceLesson `json:"Ora"`
	TeacherName    string        `json:"RogzitoTanarNeve"`
	Subject        Subject       `json:"Tantargy"`
	Type           AbsenceType   `json:"Tipus"`
	ClassGroup     ClassGroup    `json:"OsztalyCsoport"`
	ID             string        `json:"Uid"`
}

type LessonStatus struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type LessonClassGroup struct {
	ID    string `json:"Uid"`
	Title string `json:"Nev"`
}

type LessonStudentPresence struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type LessonTopic struct {
	ID    string `json:"Uid"`
	Desc  string `json:"Leiras"`
	Title string `json:"Nev"`
}

type Lesson struct {
	Status                 LessonStatus          `json:"Allapot"`
	TestIDs                []string              `json:"BejelentettSzamonkeresUids"`
	TestID                 []string              `json:"BejelentettSzamonkeresUid"`
	Date                   time.Time             `json:"Datum"`
	ViceTeacherName        string                `json:"HelyettesTanarNeve"`
	StudentHomeworkEnabled bool                  `json:"IsTanuloHazifeladatEnabled"`
	From                   time.Time             `json:"KezdetIdopont"`
	Title                  string                `json:"Nev"`
	Hours                  int                   `json:"Oraszam"`
	YearlyHours            int                   `json:"OraEvesSorszama"`
	ClassGroup             ClassGroup            `json:"OsztalyCsoport"`
	HomeworkID             string                `json:"HaziFeladatUid"`
	HomeworkSolved         bool                  `json:"IsHaziFeladatMegoldva"`
	TeacherName            string                `json:"TanarNeve"`
	Subject                Subject               `json:"Tantargy"`
	StudentPresence        LessonStudentPresence `json:"TanuloJelenlet"`
	Topic                  LessonTopic           `json:"Tema"`
	ID                     string                `json:"Uid"`
	To                     time.Time             `json:"VegIdopont"`
}

type Guardian struct {
	Email           string `json:"EmailCim"`
	Name            string `json:"Nev"`
	TelephoneNumber string `json:"Telefonszam"`
	Representative  bool   `json:"IsTorvenyesKepviselo"`
	ID              string `json:"Uid"`
}

type SchoolSystemModule struct {
	Active bool   `json:"IsAktiv"`
	Type   string `json:"Tipus"`
}

type SchoolSettings struct {
	CanStudentAssignHomework bool `json:"IsDiakRogzithetHaziFeladatot"`
	LessonTopicVisible       bool `json:"IsTanorakTemajaMegtekinthetoEllenorzoben"`
	ClassAveragesVisible     bool `json:"IsOsztalyAtlagMegjeleniteseEllenorzoben"`
	GradeVisibilityDelay     int  `json:"ErtekelesekMegjelenitesenekKesleltetesenekMerteke"`
}

type School struct {
	ID            string               `json:"Uid"`
	ShortName     string               `json:"RovidNev"`
	SystemModules []SchoolSystemModule `json:"Rendszermodulok"`
	Settings      SchoolSettings       `json:"TestreszabasBeallitasok"`
}

type UserData struct {
	MothersName string     `json:"AnyjaNeve"`
	Addresses   []string   `json:"Cimek"`
	Guardians   []Guardian `json:"Gondviselok"`
	SchoolID    string     `json:"IntezmenyAzonosito"`
	SchoolName  string     `json:"IntezmenyNev"`
	Name        string     `json:"Nev"`
	BirthDate   time.Time  `json:"SzuletesiDatum"`
	BirthPlace  string     `json:"SzuletesiHely"`
	BirthName   string     `json:"SzuletesiNev"`
	YearID      string     `json:"TanevUid"`
	ID          string     `json:"Uid"`
	School      School     `json:"Intezmeny"`
}

type Note struct{}

type Notices struct{}

type Tests struct{}

type Groups struct{}

type YearOrder struct{}

type Homework struct{}

type HomeworkComments struct{}

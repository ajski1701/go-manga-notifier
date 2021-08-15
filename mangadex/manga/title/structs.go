package title

type MangaOutput struct {
	Result string
	Data   Data
}

type Data struct {
	Id         int
	Type       int
	Attributes Attributes
}

type Attributes struct {
	Title map[string]string
}

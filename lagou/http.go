package main

type LessonsResp struct {
	State     int         `json:"state"`
	Message   string      `json:"message"`
	Content   Content     `json:"content"`
	UIMessage interface{} `json:"uiMessage"`
}
type VideoMediaDTO struct {
	ID                int         `json:"id"`
	Channel           int         `json:"channel"`
	MediaType         int         `json:"mediaType"`
	CoverImageURL     string      `json:"coverImageUrl"`
	Duration          string      `json:"duration"`
	DurationNum       int         `json:"durationNum"`
	FileID            interface{} `json:"fileId"`
	FileURL           interface{} `json:"fileUrl"`
	FileEdk           interface{} `json:"fileEdk"`
	FileSize          float64     `json:"fileSize"`
	EncryptedFileID   string      `json:"encryptedFileId"`
	MediaID           int         `json:"mediaId"`
	Clarity           int         `json:"clarity"`
	TcEncryptedFileID string      `json:"tcEncryptedFileId"`
	TcAppID           string      `json:"tcAppId"`
	TcPlayerToken     string      `json:"tcPlayerToken"`
	Free              bool        `json:"free"`
}
type AudioMediaDTO struct {
	ID                int         `json:"id"`
	Channel           int         `json:"channel"`
	MediaType         int         `json:"mediaType"`
	CoverImageURL     interface{} `json:"coverImageUrl"`
	Duration          string      `json:"duration"`
	DurationNum       int         `json:"durationNum"`
	FileID            string      `json:"fileId"`
	FileURL           string      `json:"fileUrl"`
	FileEdk           string      `json:"fileEdk"`
	FileSize          float64     `json:"fileSize"`
	EncryptedFileID   string      `json:"encryptedFileId"`
	MediaID           int         `json:"mediaId"`
	Clarity           int         `json:"clarity"`
	TcEncryptedFileID interface{} `json:"tcEncryptedFileId"`
	TcAppID           interface{} `json:"tcAppId"`
	TcPlayerToken     interface{} `json:"tcPlayerToken"`
	Free              bool        `json:"free"`
}
type CourseLessons struct {
	ID            int           `json:"id"`
	CourseID      int           `json:"courseId"`
	SectionID     int           `json:"sectionId"`
	Theme         string        `json:"theme"`
	Duration      interface{}   `json:"duration"`
	CanPlay       bool          `json:"canPlay"`
	AppID         string        `json:"appId"`
	LessonSortNum int           `json:"lessonSortNum"`
	VideoMediaDTO VideoMediaDTO `json:"videoMediaDTO"`
	AudioMediaDTO AudioMediaDTO `json:"audioMediaDTO"`
	TextContent   interface{}   `json:"textContent"`
	HandoutsURL   interface{}   `json:"handoutsUrl"`
	Status        string        `json:"status"`
	HasVideo      bool          `json:"hasVideo"`
	HasLearned    bool          `json:"hasLearned"`
	TextURL       string        `json:"textUrl"`
	PlayTypeCode  int           `json:"playTypeCode"`
	GuideCourse   bool          `json:"guideCourse"`
}
type CourseSectionList struct {
	ID             int             `json:"id"`
	CourseID       int             `json:"courseId"`
	SectionName    string          `json:"sectionName"`
	SectionSortNum int             `json:"sectionSortNum"`
	Description    string          `json:"description"`
	CourseLessons  []CourseLessons `json:"courseLessons"`
	Visible        bool            `json:"visible"`
}
type ShareInfo struct {
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ShareURL string `json:"shareUrl"`
	ImageURL string `json:"imageUrl"`
}
type Content struct {
	VideoChannelCode  int                 `json:"videoChannelCode"`
	HasBuy            bool                `json:"hasBuy"`
	CourseName        string              `json:"courseName"`
	CoverImage        string              `json:"coverImage"`
	CourseSectionList []CourseSectionList `json:"courseSectionList"`
	ShareInfo         ShareInfo           `json:"shareInfo"`
	CanHideFast       bool                `json:"canHideFast"`
	ShowHandouts      bool                `json:"showHandouts"`
	ActivityBtn       interface{}         `json:"activityBtn"`
}
type LessonResp struct {
	State     int         `json:"state"`
	Message   string      `json:"message"`
	Content   Info        `json:"content"`
	UIMessage interface{} `json:"uiMessage"`
}
type VideoMedia struct {
	ID          int         `json:"id"`
	Channel     int         `json:"channel"`
	MediaType   int         `json:"mediaType"`
	FileID      string      `json:"fileId"`
	FileURL     interface{} `json:"fileUrl"`
	Duration    string      `json:"duration"`
	FileSize    float64     `json:"fileSize"`
	DurationNum int         `json:"durationNum"`
}
type AudioMedia struct {
	ID          int     `json:"id"`
	Channel     int     `json:"channel"`
	MediaType   int     `json:"mediaType"`
	FileID      string  `json:"fileId"`
	FileURL     string  `json:"fileUrl"`
	Duration    string  `json:"duration"`
	FileSize    float64 `json:"fileSize"`
	DurationNum int     `json:"durationNum"`
}
type TeacherDTOList struct {
	ID                int    `json:"id"`
	TeacherName       string `json:"teacherName"`
	Position          string `json:"position"`
	TeacherHeadPicURL string `json:"teacherHeadPicUrl"`
	Description       string `json:"description"`
}
type Info struct {
	ID             int              `json:"id"`
	CourseID       int              `json:"courseId"`
	SectionID      int              `json:"sectionId"`
	Theme          string           `json:"theme"`
	IsFree         bool             `json:"isFree"`
	PublishDate    string           `json:"publishDate"`
	VideoMedia     VideoMedia       `json:"videoMedia"`
	AudioMedia     AudioMedia       `json:"audioMedia"`
	TeacherDTOList []TeacherDTOList `json:"teacherDTOList"`
	TextContent    string           `json:"textContent"`
	Status         string           `json:"status"`
}

type CourseResp struct {
	State     int           `json:"state"`
	Message   string        `json:"message"`
	Content   CourseContent `json:"content"`
	UIMessage interface{}   `json:"uiMessage"`
}
type MemberAdsBar struct {
	IdentityCode int    `json:"identityCode"`
	Title        string `json:"title"`
	Tips         string `json:"tips"`
	ButtonText   string `json:"buttonText"`
	URL          string `json:"url"`
	MemberType   int    `json:"memberType"`
}
type CourseRecordList struct {
	ID                    int         `json:"id"`
	Name                  string      `json:"name"`
	H5URL                 string      `json:"h5Url"`
	LastLearnLessonName   string      `json:"lastLearnLessonName"`
	Image                 string      `json:"image"`
	OpenWebToStudy        string      `json:"openWebToStudy"`
	HasCourseWeChatGroup  bool        `json:"hasCourseWeChatGroup"`
	CourseWeChatGroupTips string      `json:"courseWeChatGroupTips"`
	UpdateProgress        string      `json:"updateProgress"`
	UpdateTips            interface{} `json:"updateTips"`
	LessonUpdateNum       int         `json:"lessonUpdateNum"`
	IsEnterpriseCourse    bool        `json:"isEnterpriseCourse"`
	IsCampusCourse        bool        `json:"isCampusCourse"`
	VipFreeCourse         bool        `json:"vipFreeCourse"`
	ClassCourseType       int         `json:"classCourseType"`
}
type AllCoursePurchasedRecord struct {
	CourseType          int                `json:"courseType"`
	Title               string             `json:"title"`
	BigCourseRecordList interface{}        `json:"bigCourseRecordList"`
	CourseRecordList    []CourseRecordList `json:"courseRecordList"`
}
type CourseContent struct {
	MemberAdsBar             MemberAdsBar               `json:"memberAdsBar"`
	AllCoursePurchasedRecord []AllCoursePurchasedRecord `json:"allCoursePurchasedRecord"`
	CourseOrderSynEntry      interface{}                `json:"courseOrderSynEntry"`
	ArchivesTip              interface{}                `json:"archivesTip"`
}

package setcachepara

/*`{

"service_url"	:"testmap:MapServer",
"levels"	:"500000;250000;125000;64000;32000;16000",
"thread_count":	2,
"update_mode":	"RECREATE_ALL_TILES",
"constraining_extent":	"DEFAULT",
"area_of_interest":	{ "displayFieldName":+"", +"geometryType":+"esriGeometryPolygon", +"spatialReference":+{ ++"wkid":+4326, ++"latestWkid":+4326 +}, +"fields":+[ ++{ +++"name":+"OID", +++"type":+"esriFieldTypeOID", +++"alias":+"OID" ++}, ++{ +++"name":+"updateGeom_Length", +++"type":+"esriFieldTypeDouble", +++"alias":+"updateGeom_Length" ++}, ++{ +++"name":+"updateGeom_Area", +++"type":+"esriFieldTypeDouble", +++"alias":+"updateGeom_Area" ++} +], +"features":+[], +"exceededTransferLimit":+false },

"returnZ"	:false
"returnM":	false
"f"	:"pjson"}`*/
type parameter struct {
	ServiceURL         string `json:"service_url"`
	Levels             string `json:"levels"`
	ThreadCunt         string `json:"thread_count"`
	UpdateMode         string `json:"update_mode"`
	ConstrainingExtent string `json:"constraining_extent"`
	AreaOfInterest     string `json:"area_of_interest"`
	ReturnZ            string `json:"returnZ"`
	ReturnM            string `json:"returnM"`
	Format             string `json:"f"`
}

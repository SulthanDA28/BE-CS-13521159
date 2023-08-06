package main

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
)

type fakultas struct {
    IDfakultas int `json:"idfakultas"`
    Namafakultas string `json:"namafakultas"`
}

type jurusan struct {
    IDjurusan int `json:"idjurusan"`
    Namajurusan string `json:"namajurusan"`
    FakultasID int `json:"fakultas_id"`
}

type matkulfakultas struct {
    IDmatkul int `json:"idmatkul"`
    Namamatkul string `json:"namamatkul"`
    FakultasID int `json:"fakultas_id"`
    Sks int `json:"sks"`
    Minsemester int `json:"minsemester"`
    Prediksinilai string `json:"prediksinilai"`
}

type matkuljurusan struct {
    IDmatkul int `json:"idmatkul"`
    Namamatkul string `json:"namamatkul"`
    JurusanID int `json:"jurusan_id"`
    Sks int `json:"sks"`
    Minsemester int `json:"minsemester"`
    Prediksinilai string `json:"prediksinilai"`
}

func getFakultas() []fakultas {
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM fakultas")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var result []fakultas
    for rows.Next() {
        var each = fakultas{}
        var err = rows.Scan(&each.IDfakultas, &each.Namafakultas)
        if err != nil {
            log.Fatal(err)
        }
        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    return result

}

func getJurusan() []jurusan {
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM jurusan")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var result []jurusan
    for rows.Next() {
        var each = jurusan{}
        var err = rows.Scan(&each.IDjurusan, &each.Namajurusan, &each.FakultasID)
        if err != nil {
            log.Fatal(err)
        }
        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    return result

}

func getMatkulFakultas() []matkulfakultas {
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM matkulfakultas")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var result []matkulfakultas
    for rows.Next() {
        var each = matkulfakultas{}
        var err = rows.Scan(&each.IDmatkul, &each.Namamatkul, &each.Sks, &each.Minsemester,&each.FakultasID, &each.Prediksinilai)
        if err != nil {
            log.Fatal(err)
        }
        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    return result

}

func getMatkulJurusan() []matkuljurusan {
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM matkuljurusan")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var result []matkuljurusan
    for rows.Next() {
        var each = matkuljurusan{}
        var err = rows.Scan(&each.IDmatkul, &each.Namamatkul, &each.Sks, &each.Minsemester,&each.JurusanID, &each.Prediksinilai)
        if err != nil {
            log.Fatal(err)
        }
        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    return result

}

func insertFakultas(namafakultas string) string {
    ceknamafakultas := strings.ToLower(namafakultas)
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    ambilfakultas := getFakultas()
    for _, each := range ambilfakultas {
        each.Namafakultas = strings.ToLower(each.Namafakultas)
        if each.Namafakultas == ceknamafakultas {
            namafakultas = ""
        }
    }
    if (namafakultas == ""){
        return "fakultas sudah ada"
    }else{
        _, err = db.Exec("INSERT INTO fakultas (nama) VALUES (?)", namafakultas)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil ditambahkan dengan fakultas: "+namafakultas
    }
}

func insertJurusan(namajurusan string, namafakultas string) string{
    ceknamajurusan := strings.ToLower(namajurusan)
    namafakultasasli := namafakultas
    namafakultas = strings.ToLower(namafakultas)
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    ambilfakultas := getFakultas()
    var idfakultas int
    for _, each := range ambilfakultas {
        each.Namafakultas = strings.ToLower(each.Namafakultas)
        if each.Namafakultas == namafakultas {
            idfakultas = each.IDfakultas
        }
    }
    ambiljurusan := getJurusan()
    sama := false
    for _, each := range ambiljurusan {
        each.Namajurusan = strings.ToLower(each.Namajurusan)
        if each.Namajurusan == ceknamajurusan {
            sama = true
            break
        }
    }
    if idfakultas == 0 {
        return "fakultas tidak ditemukan"
    }else if (sama){
        _, err = db.Exec("UPDATE jurusan SET fakultas_id = ? WHERE nama = ?", idfakultas, namajurusan)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil diupdate pada "+namajurusan+":"+namafakultasasli
    }else{
        _, err = db.Exec("INSERT INTO jurusan (nama, fakultas_id) VALUES (?, ?)", namajurusan, idfakultas)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil ditambahkan pada "+namajurusan+":"+namafakultasasli
    }

}

func insertMatkulFakultas(namamatkul string, namafakultas string, sks int, minsemester int, prediksinilai string) string{
    ceknamamatkul := strings.ToLower(namamatkul)
    namafakultasasli := namafakultas
    namafakultas = strings.ToLower(namafakultas)
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    ambilfakultas := getFakultas()
    var idfakultas int
    for _, each := range ambilfakultas {
        each.Namafakultas = strings.ToLower(each.Namafakultas)
        if each.Namafakultas == namafakultas {
            idfakultas = each.IDfakultas
        }
    }
    ambilmatkulfakultas := getMatkulFakultas()
    sama := false
    for _, each := range ambilmatkulfakultas {
        each.Namamatkul = strings.ToLower(each.Namamatkul)
        if each.Namamatkul == ceknamamatkul && each.FakultasID == idfakultas {
            sama = true
            break
        }
    }
    if idfakultas == 0 {
        return "fakultas tidak ditemukan"
    }else if (sama){
        _, err = db.Exec("UPDATE matkulfakultas SET sks = ?, semestermin = ?, prediksinilai = ? WHERE nama = ? and fakultas_id = ?", sks, minsemester, prediksinilai, namamatkul, idfakultas)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil diupdate pada "+namamatkul+":"+namafakultasasli
    }else{
        _, err = db.Exec("INSERT INTO matkulfakultas (nama, fakultas_id, sks, semestermin, prediksinilai) VALUES (?, ?, ?, ?, ?)", namamatkul, idfakultas, sks, minsemester, prediksinilai)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil ditambahkan pada "+namamatkul+":"+namafakultasasli
    }

}

func insertMatkulJurusan(namamatkul string, namajurusan string, sks int, minsemester int, prediksinilai string) string{
    ceknamamatkul := strings.ToLower(namamatkul)
    namajurusanasli := namajurusan
    namajurusan = strings.ToLower(namajurusan)
    db, err := sql.Open("mysql", "username:password@tcp(database:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    ambiljurusan := getJurusan()
    var idjurusan int
    for _, each := range ambiljurusan {
        each.Namajurusan = strings.ToLower(each.Namajurusan)
        if each.Namajurusan == namajurusan {
            idjurusan = each.IDjurusan
        }
    }
    ambilmatkuljurusan := getMatkulJurusan()
    sama := false
    for _, each := range ambilmatkuljurusan {
        each.Namamatkul = strings.ToLower(each.Namamatkul)
        if each.Namamatkul == ceknamamatkul && each.JurusanID == idjurusan {
            sama = true
            break
        }
    }
    if idjurusan == 0 {
        return "jurusan tidak ditemukan"
    }else if (sama){
        _, err = db.Exec("UPDATE matkuljurusan SET sks = ?, semestermin = ?, prediksinilai = ? WHERE nama = ? and jurusan_id = ?", sks, minsemester, prediksinilai, namamatkul, idjurusan)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil diupdate pada "+namamatkul+":"+namajurusanasli
    }else{
        _, err = db.Exec("INSERT INTO matkuljurusan (nama, jurusan_id, sks, semestermin, prediksinilai) VALUES (?, ?, ?, ?, ?)", namamatkul, idjurusan, sks, minsemester, prediksinilai)
        if err != nil {
            log.Fatal(err)
        }
        return "Data berhasil ditambahkan pada "+namamatkul+":"+namajurusanasli
    }

}

type buattampilan struct {
    Tipe string `json:"tipe"`
    Nama string `json:"nama"`
    Namamatkul string `json:"namamatkul"`
    Sks int `json:"sks"`
    Minsemester int `json:"minsemester"`
    Prediksinilai string `json:"prediksinilai"`
}

func sendMatkulFakultas(namafakultas string,ambilsemester int) []buattampilan{
    namaasli := namafakultas
    namafakultas = strings.ToLower(namafakultas)
    ambilfakultas := getFakultas()
    var idfakultas int
    for _, each := range ambilfakultas {
        each.Namafakultas = strings.ToLower(each.Namafakultas)
        if each.Namafakultas == namafakultas {
            idfakultas = each.IDfakultas
        }
    }
    simpanmatkulfakultas := []buattampilan{}
    ambilmatkulfakultas := getMatkulFakultas()
    for _, each := range ambilmatkulfakultas {
        if each.FakultasID == idfakultas {
            if (ambilsemester>=each.Minsemester){
                var tampung buattampilan
                tampung.Tipe = "fakultas"
                tampung.Nama = namaasli
                tampung.Namamatkul = each.Namamatkul
                tampung.Sks = each.Sks
                tampung.Minsemester = each.Minsemester
                tampung.Prediksinilai = each.Prediksinilai
                simpanmatkulfakultas = append(simpanmatkulfakultas, tampung)
            }
        }
    }
    return simpanmatkulfakultas

}


func sendMatkulJurusan(namajurusan string,ambilsemester int) []buattampilan{
    namajurusanasli := namajurusan
    namajurusan = strings.ToLower(namajurusan)

    ambiljurusan := getJurusan()
    var idjurusan int
    var buatcekfakultas jurusan
    for _, each := range ambiljurusan {
        each.Namajurusan = strings.ToLower(each.Namajurusan)
        if each.Namajurusan == namajurusan {
            idjurusan = each.IDjurusan
            buatcekfakultas = each
        }
    }
    simpanmatkuljurusan := []buattampilan{}
    ambilmatkuljurusan := getMatkulJurusan()
    for _, each := range ambilmatkuljurusan {
        if each.JurusanID == idjurusan {
            if (ambilsemester>=each.Minsemester){
                var tampung buattampilan
                tampung.Tipe = "jurusan"
                tampung.Nama = namajurusanasli
                tampung.Namamatkul = each.Namamatkul
                tampung.Sks = each.Sks
                tampung.Minsemester = each.Minsemester
                tampung.Prediksinilai = each.Prediksinilai
                simpanmatkuljurusan = append(simpanmatkuljurusan, tampung)
            }
        }
    }
    ambilnamafakultas := getFakultas()
    var fakultasnama string
    for _, each := range ambilnamafakultas {
        if each.IDfakultas == buatcekfakultas.FakultasID {
            fakultasnama = each.Namafakultas
        }
    }

    cekmatkulfakultas := sendMatkulFakultas(fakultasnama,ambilsemester)
    for _, each := range cekmatkulfakultas {
            var tampung buattampilan
            tampung.Tipe = "fakultas"
            tampung.Nama = fakultasnama
            tampung.Namamatkul = each.Namamatkul
            tampung.Sks = each.Sks
            tampung.Minsemester = each.Minsemester
            tampung.Prediksinilai = each.Prediksinilai
            simpanmatkuljurusan = append(simpanmatkuljurusan, tampung)
    }
    return simpanmatkuljurusan

}


//-------------------
//API
//-------------------

func apiGetFakultas(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
    var result = getFakultas()
    c.JSON(http.StatusOK, result)
}
func apiGetJurusan(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
    var result = getJurusan()
    
    c.JSON(http.StatusOK, result)
}
type fakultasInput struct {
    Namafakultas string `json:"namafakultas" binding:"required"`
}
func apiPostFakultas(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
    var input fakultasInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hasil:=insertFakultas(input.Namafakultas)
    c.JSON(http.StatusOK, gin.H{"status": hasil})
}
type jurusanInput struct {
    Namajurusan string `json:"namajurusan" binding:"required"`
    NamaFakultas string `json:"namafakultas" binding:"required"`
}
func apiPostJurusan(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
    var input jurusanInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hasil := insertJurusan(input.Namajurusan, input.NamaFakultas)
    c.JSON(http.StatusOK, gin.H{"status": hasil})
}

type matkulFakultasInput struct {
    Namamatkul string `json:"namamatkul" binding:"required"`
    NamaFakultas string `json:"namafakultas" binding:"required"`
    Sks int `json:"sks" binding:"required"`
    Minsemester int `json:"minsemester" binding:"required"`
    Prediksinilai string `json:"prediksinilai" binding:"required"`
}

func apiPostMatkulFakultas(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input matkulFakultasInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hasil := insertMatkulFakultas(input.Namamatkul, input.NamaFakultas, input.Sks, input.Minsemester, input.Prediksinilai)
    c.JSON(http.StatusOK, gin.H{"status": hasil})
}

type matkulJurusanInput struct {
    Namamatkul string `json:"namamatkul" binding:"required"`
    NamaJurusan string `json:"namajurusan" binding:"required"`
    Sks int `json:"sks" binding:"required"`
    Minsemester int `json:"minsemester" binding:"required"`
    Prediksinilai string `json:"prediksinilai" binding:"required"`
}

func apiPostMatkulJurusan(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input matkulJurusanInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hasil := insertMatkulJurusan(input.Namamatkul, input.NamaJurusan, input.Sks, input.Minsemester, input.Prediksinilai)
    c.JSON(http.StatusOK, gin.H{"status": hasil})
}

func apiGetMatkulFakultas(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var result = getMatkulFakultas()
    c.JSON(http.StatusOK, result)
}

func apiGetMatkulJurusan(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var result = getMatkulJurusan()
    c.JSON(http.StatusOK, result)
}

type fakultasInputNama struct {
    Namafakultas string `json:"namafakultas" binding:"required"`
    Ambilsemester int `json:"ambilsemester" binding:"required"`
}

func apiGetMatkulFakultasNama(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input fakultasInputNama
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var result = sendMatkulFakultas(input.Namafakultas,input.Ambilsemester)
    c.JSON(http.StatusOK, result)
}

type jurusanInputNama struct {
    Namajurusan string `json:"namajurusan" binding:"required"`
    Ambilsemester int `json:"ambilsemester" binding:"required"`
}

func apiGetMatkulJurusanNama(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input jurusanInputNama
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var result = sendMatkulJurusan(input.Namajurusan,input.Ambilsemester)
    c.JSON(http.StatusOK, result)
}

type fakultasInputPrediksi struct {
    Namafakultas string `json:"namafakultas" binding:"required"`
    Minsemester int `json:"minsemester" binding:"required"`
    Maxsks int `json:"maxsks" binding:"required"`
    Minsks int `json:"minsks" binding:"required"`
}

func apiGetPrediksiFakultas(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input fakultasInputPrediksi
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var result = prediksinilaiFakultas(input.Namafakultas,input.Minsemester,input.Maxsks,input.Minsks)
    c.JSON(http.StatusOK, result)
}

type jurusanInputPrediksi struct {
    Namajurusan string `json:"namajurusan" binding:"required"`
    Minsemester int `json:"minsemester" binding:"required"`
    Maxsks int `json:"maxsks" binding:"required"`
    Minsks int `json:"minsks" binding:"required"`
}

func apiGetPrediksiJurusan(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    var input jurusanInputPrediksi
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var result = prediksinilaiJurusan(input.Namajurusan,input.Minsemester,input.Maxsks,input.Minsks)
    c.JSON(http.StatusOK, result)
}


//--------------------------
//---Dynamic Programming----
//--------------------------

func prediksinilaimatkul(sks []int,prednilai []string,max int,min int)(float32,[]int,int){//dynamic programming
    var skstofloat []float32//bobot
    n := len(sks)
    for _, each := range sks {
        skstofloat = append(skstofloat, float32(each))
    }
    var convert []float32//value
    for i, each := range prednilai {
        if each == "A" {
            convert = append(convert, 4*skstofloat[i])
        }else if each == "AB" {
            convert = append(convert, 3.5*skstofloat[i])
        }else if each == "B" {
            convert = append(convert, 3*skstofloat[i])
        }else if each == "BC" {
            convert = append(convert, 2.5*skstofloat[i])
        }else if each == "C" {
            convert = append(convert, 2*skstofloat[i])
        }else if each == "D" {
            convert = append(convert, 1*skstofloat[i])
        }else if each == "E" {
            convert = append(convert, 0*skstofloat[i])
        }
    }
    var totalnilai float32
    for _, each := range convert {
        totalnilai += each
    }
    simpan := make([][][]int, n+1)
    for i := range simpan {
        simpan[i] = make([][]int, max+1)
    }
    dp := make([][]float32, n+1)
    for i := range dp {
        dp[i] = make([]float32, max+1)
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= max; j++ {
            cek := j - sks[i-1]
            if (cek<0){
                dp[i][j] = dp[i-1][j]
                simpan[i][j] = simpan[i-1][j]
            }else{
                if (dp[i-1][j] > dp[i-1][cek]+convert[i-1]) {
                    dp[i][j] = dp[i-1][j]
                    simpan[i][j] = simpan[i-1][j]
                }else{
                    dp[i][j] = dp[i-1][cek]+convert[i-1]
                    simpan[i][j] = append(simpan[i-1][cek], i-1)
                }
            }

        }
    }
    hasilnilaibawah := dp[n]
    indeksbawah := simpan[n]
    var indeks int
    var maxnilai float32
    for i := 0; i < len(hasilnilaibawah); i++ {
        cekminbobot := 0
        for _, each := range indeksbawah[i] {
            cekminbobot += sks[each]
        }
        if (hasilnilaibawah[i] > maxnilai && cekminbobot >= min) {
            maxnilai = hasilnilaibawah[i]
            indeks = i
        }
    }
    totalsks := 0
    for _, each := range indeksbawah[indeks] {
        totalsks += sks[each]
    }
    ipk := maxnilai/float32(totalsks)

    
    return ipk,indeksbawah[indeks],totalsks
}

type hasilakhirbanget struct {
    Ipk float32 `json:"ipk"`
    Totalsks int `json:"totalsks"`
    Semuamatkul []buattampilan `json:"semuamatkul"`
}

func prediksinilaiFakultas(namafakultas string,minsemester int,maxsks int,minsks int) hasilakhirbanget{
    getallmatkulfakultas := sendMatkulFakultas(namafakultas,minsemester)
    var sks []int
    var prednilai []string
    for _, each := range getallmatkulfakultas {
        sks = append(sks, each.Sks)
        prednilai = append(prednilai, each.Prediksinilai)
    }
    ipk,indeksbawah,totalsks := prediksinilaimatkul(sks,prednilai,maxsks,minsks)
    var semuamatkul []buattampilan
    for _, each := range indeksbawah {
        semuamatkul = append(semuamatkul, getallmatkulfakultas[each])
    }
    var hasilakhir hasilakhirbanget
    if (semuamatkul == nil){
        var semuamatkul []buattampilan
        hasilakhir.Ipk = 0
        hasilakhir.Totalsks = 0
        hasilakhir.Semuamatkul = semuamatkul
        return hasilakhir
    }else{
        hasilakhir.Ipk = ipk
        hasilakhir.Totalsks = totalsks
        hasilakhir.Semuamatkul = semuamatkul
    }
    return hasilakhir

}

func prediksinilaiJurusan(namajurusan string,minsemester int,maxsks int,minsks int) hasilakhirbanget{
    getallmatkuljurusan := sendMatkulJurusan(namajurusan,minsemester)
    var sks []int
    var prednilai []string
    for _, each := range getallmatkuljurusan {
        sks = append(sks, each.Sks)
        prednilai = append(prednilai, each.Prediksinilai)
    }
    ipk,indeksbawah,totalsks := prediksinilaimatkul(sks,prednilai,maxsks,minsks)
    var semuamatkul []buattampilan
    for _, each := range indeksbawah {
        semuamatkul = append(semuamatkul, getallmatkuljurusan[each])
    }
    var hasilakhir hasilakhirbanget
    if (semuamatkul == nil){
        var semuamatkul []buattampilan
        hasilakhir.Ipk = 0
        hasilakhir.Totalsks = 0
        hasilakhir.Semuamatkul = semuamatkul
        return hasilakhir
    }else{
        hasilakhir.Ipk = ipk
        hasilakhir.Totalsks = totalsks
        hasilakhir.Semuamatkul = semuamatkul
        return hasilakhir
    }

}







func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World!")
    })
    router.GET("/fakultas", apiGetFakultas)
    router.GET("/jurusan", apiGetJurusan)
    router.POST("/fakultasadd", apiPostFakultas)
    router.POST("/jurusanadd", apiPostJurusan)
    router.POST("/matkulfakultasadd", apiPostMatkulFakultas)
    router.POST("/matkuljurusanadd", apiPostMatkulJurusan)
    router.GET("/matkulfakultas", apiGetMatkulFakultas)
    router.GET("/matkuljurusan", apiGetMatkulJurusan)
    router.POST("/matkulfakultasnama", apiGetMatkulFakultasNama)
    router.POST("/matkuljurusannama", apiGetMatkulJurusanNama)
    router.POST("/prediksifakultas", apiGetPrediksiFakultas)
    router.POST("/prediksijurusan", apiGetPrediksiJurusan)
    router.Use(cors.Default())
    router.Run(":8080")

}

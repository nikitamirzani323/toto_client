<script>
  import {
    Button,
    Card,
    CardBody,
    CardHeader,
    TabContent,
    TabPane,
  } from "sveltestrap";
  import Modal from "../components/Modal.svelte";
  import Tablekeranjang5050 from "../permainan/Tablekeranjang5050.svelte";
  import Loader from "../components/Loader.svelte";
  import { createEventDispatcher } from "svelte";
  import { notifications } from "../components/Noti.svelte";
  import PeriodePanel from "../components/PeriodePanel.svelte";
  import Swal from "sweetalert2";

  export let idcomppasaran = "";
  export let idtrxkeluaran = "";
  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_ipaddress = "";
  export let client_device = "";
  export let client_timezone = "";
  export let pasaran_code = "";
  export let pasaran_name = "";
  export let pasaran_periode = 0;
  export let permainan_title = "COLOK";
  export let daylight = false;

  const swalWithBootstrapButtons = Swal.mixin({
    customClass: {
      confirmButton: "mx-2 rounded-2 btn btn-success",
      cancelButton: "mx-2 rounded-2 btn btn-danger",
    },
    buttonsStyling: false,
  });

  let keranjang = [];
  let css_loader = "display:none;";
  let nomor_global = 0;
  let totalkeranjang = 0;
  let group_btn_beli = false;
  let record = "";
  let temp_bulk_error = "";

  let min_bet_5050umum = 0;
  let max_bet_5050umum = 0;
  let keibesar_bet_5050umum = 0;
  let keikecil_bet_5050umum = 0;
  let keigenap_bet_5050umum = 0;
  let keiganjil_bet_5050umum = 0;
  let keitengah_bet_5050umum = 0;
  let keitepi_bet_5050umum = 0;
  let discbesar_bet_5050umum = 0;
  let disckecil_bet_5050umum = 0;
  let discgenap_bet_5050umum = 0;
  let discganjil_bet_5050umum = 0;
  let disctengah_bet_5050umum = 0;
  let disctepi_bet_5050umum = 0;
  let limittotal_bet_5050umum = 0;
  let min_bet_5050special = 0;
  let max_bet_5050special = 0;
  let keiasganjil_bet_5050special = 0;
  let keiasgenap_bet_5050special = 0;
  let keiasbesar_bet_5050special = 0;
  let keiaskecil_bet_5050special = 0;
  let keikopganjil_bet_5050special = 0;
  let keikopgenap_bet_5050special = 0;
  let keikopbesar_bet_5050special = 0;
  let keikopkecil_bet_5050special = 0;
  let keikepalaganjil_bet_5050special = 0;
  let keikepalagenap_bet_5050special = 0;
  let keikepalabesar_bet_5050special = 0;
  let keikepalakecil_bet_5050special = 0;
  let keiekorganjil_bet_5050special = 0;
  let keiekorgenap_bet_5050special = 0;
  let keiekorbesar_bet_5050special = 0;
  let keiekorkecil_bet_5050special = 0;
  let discasganjil_bet_5050special = 0;
  let discasgenap_bet_5050special = 0;
  let discasbesar_bet_5050special = 0;
  let discaskecil_bet_5050special = 0;
  let disckopganjil_bet_5050special = 0;
  let disckopgenap_bet_5050special = 0;
  let disckopbesar_bet_5050special = 0;
  let disckopkecil_bet_5050special = 0;
  let disckepalaganjil_bet_5050special = 0;
  let disckepalagenap_bet_5050special = 0;
  let disckepalabesar_bet_5050special = 0;
  let disckepalakecil_bet_5050special = 0;
  let discekorganjil_bet_5050special = 0;
  let discekorgenap_bet_5050special = 0;
  let discekorbesar_bet_5050special = 0;
  let discekorkecil_bet_5050special = 0;
  let limittotal_bet_5050special = 0;
  let min_bet_5050kombinasi = 0;
  let max_bet_5050kombinasi = 0;
  let kei_belakangmono_bet_5050kombinasi = 0;
  let kei_belakangstereo_bet_5050kombinasi = 0;
  let kei_belakangkembang_bet_5050kombinasi = 0;
  let kei_belakangkempis_bet_5050kombinasi = 0;
  let kei_belakangkembar_bet_5050kombinasi = 0;
  let kei_tengahmono_bet_5050kombinasi = 0;
  let kei_tengahstereo_bet_5050kombinasi = 0;
  let kei_tengahkembang_bet_5050kombinasi = 0;
  let kei_tengahkempis_bet_5050kombinasi = 0;
  let kei_tengahkembar_bet_5050kombinasi = 0;
  let kei_depanmono_bet_5050kombinasi = 0;
  let kei_depanstereo_bet_5050kombinasi = 0;
  let kei_depankembang_bet_5050kombinasi = 0;
  let kei_depankempis_bet_5050kombinasi = 0;
  let kei_depankembar_bet_5050kombinasi = 0;
  let disc_belakangmono_bet_5050kombinasi = 0;
  let disc_belakangstereo_bet_5050kombinasi = 0;
  let disc_belakangkembang_bet_5050kombinasi = 0;
  let disc_belakangkempis_bet_5050kombinasi = 0;
  let disc_belakangkembar_bet_5050kombinasi = 0;
  let disc_tengahmono_bet_5050kombinasi = 0;
  let disc_tengahstereo_bet_5050kombinasi = 0;
  let disc_tengahkembang_bet_5050kombinasi = 0;
  let disc_tengahkempis_bet_5050kombinasi = 0;
  let disc_tengahkembar_bet_5050kombinasi = 0;
  let disc_depanmono_bet_5050kombinasi = 0;
  let disc_depanstereo_bet_5050kombinasi = 0;
  let disc_depankembang_bet_5050kombinasi = 0;
  let disc_depankempis_bet_5050kombinasi = 0;
  let disc_depankembar_bet_5050kombinasi = 0;
  let limittotal_bet_5050kombinasi = 0;

  let count_line_5050umum = 0;
  let count_line_5050special = 0;
  let count_line_5050kombinasi = 0;

  let db_form5050_umum_count_temp = 0;
  let db_form5050_special_count_temp = 0;
  let db_form5050_kombinasi_count_temp = 0;
  let flag_fulldiskon = "DISC";
  let dispatch = createEventDispatcher();

  async function inittogel_432d(e) {
    const res = await fetch("/api/inittogel_432d", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        pasaran_code: pasaran_code,
        permainan: e,
      }),
    });
    group_btn_beli = true;
    const json = await res.json();
    record = json.record;

    for (var i = 0; i < record.length; i++) {
      min_bet_5050umum = parseInt(record[i]["min_bet_5050umum"]);
      max_bet_5050umum = parseInt(record[i]["max_bet_5050umum"]);
      keibesar_bet_5050umum = parseFloat(record[i]["keibesar_bet_5050umum"]);
      keikecil_bet_5050umum = parseFloat(record[i]["keikecil_bet_5050umum"]);
      keigenap_bet_5050umum = parseFloat(record[i]["keigenap_bet_5050umum"]);
      keiganjil_bet_5050umum = parseFloat(record[i]["keiganjil_bet_5050umum"]);
      keitengah_bet_5050umum = parseFloat(record[i]["keitengah_bet_5050umum"]);
      keitepi_bet_5050umum = parseFloat(record[i]["keitepi_bet_5050umum"]);
      discbesar_bet_5050umum = parseFloat(record[i]["discbesar_bet_5050umum"]);
      disckecil_bet_5050umum = parseFloat(record[i]["disckecil_bet_5050umum"]);
      discgenap_bet_5050umum = parseFloat(record[i]["discgenap_bet_5050umum"]);
      discganjil_bet_5050umum = parseFloat(
        record[i]["discganjil_bet_5050umum"]
      );
      disctengah_bet_5050umum = parseFloat(
        record[i]["disctengah_bet_5050umum"]
      );
      disctepi_bet_5050umum = parseFloat(record[i]["disctepi_bet_5050umum"]);
      limittotal_bet_5050umum = parseFloat(
        record[i]["limittotal_bet_5050umum"]
      );
      min_bet_5050special = parseFloat(record[i]["min_bet_5050special"]);
      max_bet_5050special = parseFloat(record[i]["max_bet_5050special"]);
      keiasganjil_bet_5050special = parseFloat(
        record[i]["keiasganjil_bet_5050special"]
      );
      keiasgenap_bet_5050special = parseFloat(
        record[i]["keiasgenap_bet_5050special"]
      );
      keiasbesar_bet_5050special = parseFloat(
        record[i]["keiasbesar_bet_5050special"]
      );
      keiaskecil_bet_5050special = parseFloat(
        record[i]["keiaskecil_bet_5050special"]
      );
      keikopganjil_bet_5050special = parseFloat(
        record[i]["keikopganjil_bet_5050special"]
      );
      keikopgenap_bet_5050special = parseFloat(
        record[i]["keikopgenap_bet_5050special"]
      );
      keikopbesar_bet_5050special = parseFloat(
        record[i]["keikopbesar_bet_5050special"]
      );
      keikopkecil_bet_5050special = parseFloat(
        record[i]["keikopkecil_bet_5050special"]
      );
      keikepalaganjil_bet_5050special = parseFloat(
        record[i]["keikepalaganjil_bet_5050special"]
      );
      keikepalagenap_bet_5050special = parseFloat(
        record[i]["keikepalagenap_bet_5050special"]
      );
      keikepalabesar_bet_5050special = parseFloat(
        record[i]["keikepalabesar_bet_5050special"]
      );
      keikepalakecil_bet_5050special = parseFloat(
        record[i]["keikepalakecil_bet_5050special"]
      );
      keiekorganjil_bet_5050special = parseFloat(
        record[i]["keiekorganjil_bet_5050special"]
      );
      keiekorgenap_bet_5050special = parseFloat(
        record[i]["keiekorgenap_bet_5050special"]
      );
      keiekorbesar_bet_5050special = parseFloat(
        record[i]["keiekorbesar_bet_5050special"]
      );
      keiekorkecil_bet_5050special = parseFloat(
        record[i]["keiekorkecil_bet_5050special"]
      );
      discasganjil_bet_5050special = parseFloat(
        record[i]["discasganjil_bet_5050special"]
      );
      discasgenap_bet_5050special = parseFloat(
        record[i]["discasgenap_bet_5050special"]
      );
      discasbesar_bet_5050special = parseFloat(
        record[i]["discasbesar_bet_5050special"]
      );
      discaskecil_bet_5050special = parseFloat(
        record[i]["discaskecil_bet_5050special"]
      );
      disckopganjil_bet_5050special = parseFloat(
        record[i]["disckopganjil_bet_5050special"]
      );
      disckopgenap_bet_5050special = parseFloat(
        record[i]["disckopgenap_bet_5050special"]
      );
      disckopbesar_bet_5050special = parseFloat(
        record[i]["disckopbesar_bet_5050special"]
      );
      disckopkecil_bet_5050special = parseFloat(
        record[i]["disckopkecil_bet_5050special"]
      );
      disckepalaganjil_bet_5050special = parseFloat(
        record[i]["disckepalaganjil_bet_5050special"]
      );
      disckepalagenap_bet_5050special = parseFloat(
        record[i]["disckepalagenap_bet_5050special"]
      );
      disckepalabesar_bet_5050special = parseFloat(
        record[i]["disckepalabesar_bet_5050special"]
      );
      disckepalakecil_bet_5050special = parseFloat(
        record[i]["disckepalakecil_bet_5050special"]
      );
      discekorganjil_bet_5050special = parseFloat(
        record[i]["discekorganjil_bet_5050special"]
      );
      discekorgenap_bet_5050special = parseFloat(
        record[i]["discekorgenap_bet_5050special"]
      );
      discekorbesar_bet_5050special = parseFloat(
        record[i]["discekorbesar_bet_5050special"]
      );
      discekorkecil_bet_5050special = parseFloat(
        record[i]["discekorkecil_bet_5050special"]
      );
      limittotal_bet_5050special = parseFloat(
        record[i]["limittotal_bet_5050special"]
      );
      min_bet_5050kombinasi = parseFloat(record[i]["min_bet_5050kombinasi"]);
      max_bet_5050kombinasi = parseFloat(record[i]["max_bet_5050kombinasi"]);
      kei_belakangmono_bet_5050kombinasi = parseFloat(
        record[i]["kei_belakangmono_bet_5050kombinasi"]
      );
      kei_belakangstereo_bet_5050kombinasi = parseFloat(
        record[i]["kei_belakangstereo_bet_5050kombinasi"]
      );
      kei_belakangkembang_bet_5050kombinasi = parseFloat(
        record[i]["kei_belakangkembang_bet_5050kombinasi"]
      );
      kei_belakangkempis_bet_5050kombinasi = parseFloat(
        record[i]["kei_belakangkempis_bet_5050kombinasi"]
      );
      kei_belakangkembar_bet_5050kombinasi = parseFloat(
        record[i]["kei_belakangkembar_bet_5050kombinasi"]
      );
      kei_tengahmono_bet_5050kombinasi = parseFloat(
        record[i]["kei_tengahmono_bet_5050kombinasi"]
      );
      kei_tengahstereo_bet_5050kombinasi = parseFloat(
        record[i]["kei_tengahstereo_bet_5050kombinasi"]
      );
      kei_tengahkembang_bet_5050kombinasi = parseFloat(
        record[i]["kei_tengahkembang_bet_5050kombinasi"]
      );
      kei_tengahkempis_bet_5050kombinasi = parseFloat(
        record[i]["kei_tengahkempis_bet_5050kombinasi"]
      );
      kei_tengahkembar_bet_5050kombinasi = parseFloat(
        record[i]["kei_tengahkembar_bet_5050kombinasi"]
      );
      kei_depanmono_bet_5050kombinasi = parseFloat(
        record[i]["kei_depanmono_bet_5050kombinasi"]
      );
      kei_depanstereo_bet_5050kombinasi = parseFloat(
        record[i]["kei_depanstereo_bet_5050kombinasi"]
      );
      kei_depankembang_bet_5050kombinasi = parseFloat(
        record[i]["kei_depankembang_bet_5050kombinasi"]
      );
      kei_depankempis_bet_5050kombinasi = parseFloat(
        record[i]["kei_depankempis_bet_5050kombinasi"]
      );
      kei_depankembar_bet_5050kombinasi = parseFloat(
        record[i]["kei_depankembar_bet_5050kombinasi"]
      );
      disc_belakangmono_bet_5050kombinasi = parseFloat(
        record[i]["disc_belakangmono_bet_5050kombinasi"]
      );
      disc_belakangstereo_bet_5050kombinasi = parseFloat(
        record[i]["disc_belakangstereo_bet_5050kombinasi"]
      );
      disc_belakangkembang_bet_5050kombinasi = parseFloat(
        record[i]["disc_belakangkembang_bet_5050kombinasi"]
      );
      disc_belakangkempis_bet_5050kombinasi = parseFloat(
        record[i]["disc_belakangkempis_bet_5050kombinasi"]
      );
      disc_belakangkembar_bet_5050kombinasi = parseFloat(
        record[i]["disc_belakangkembar_bet_5050kombinasi"]
      );
      disc_tengahmono_bet_5050kombinasi = parseFloat(
        record[i]["disc_tengahmono_bet_5050kombinasi"]
      );
      disc_tengahstereo_bet_5050kombinasi = parseFloat(
        record[i]["disc_tengahstereo_bet_5050kombinasi"]
      );
      disc_tengahkembang_bet_5050kombinasi = parseFloat(
        record[i]["disc_tengahkembang_bet_5050kombinasi"]
      );
      disc_tengahkempis_bet_5050kombinasi = parseFloat(
        record[i]["disc_tengahkempis_bet_5050kombinasi"]
      );
      disc_tengahkembar_bet_5050kombinasi = parseFloat(
        record[i]["disc_tengahkembar_bet_5050kombinasi"]
      );
      disc_depanmono_bet_5050kombinasi = parseFloat(
        record[i]["disc_depanmono_bet_5050kombinasi"]
      );
      disc_depanstereo_bet_5050kombinasi = parseFloat(
        record[i]["disc_depanstereo_bet_5050kombinasi"]
      );
      disc_depankembang_bet_5050kombinasi = parseFloat(
        record[i]["disc_depankembang_bet_5050kombinasi"]
      );
      disc_depankempis_bet_5050kombinasi = parseFloat(
        record[i]["disc_depankempis_bet_5050kombinasi"]
      );
      disc_depankembar_bet_5050kombinasi = parseFloat(
        record[i]["disc_depankembar_bet_5050kombinasi"]
      );
      limittotal_bet_5050kombinasi = parseFloat(
        record[i]["limittotal_bet_5050kombinasi"]
      );
    }
  }
  async function savetransaksi() {
    if (client_device == "WEBSITE") {
      css_loader =
        "position:absolute;z-index: 1000;left: 50%;top: 35%;display:inline;";
    } else {
      css_loader =
        "position:absolute;z-index: 1000;left: 0%;top: 35%;display:inline;";
    }
    group_btn_beli = false;
    const res = await fetch("/api/savetransaksi", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        transaction: reverseString(
          btoa(
            JSON.stringify({
              pasaran_idtransaction: idtrxkeluaran,
              pasaran_idcomp: idcomppasaran,
              pasaran_code: pasaran_code,
              pasaran_periode: pasaran_periode,
              token: client_token,
              company: client_company,
              username: client_username,
              ipaddress: client_ipaddress,
              devicemember: client_device,
              timezone: client_timezone,
              total: totalkeranjang,
              data: keranjang,
            })
          )
        ),
      }),
    });
    const json = await res.json();
    if (json.status == "200") {
      css_loader = "display:none;";
      Swal.fire({
        position: "center",
        icon: "success",
        title: "Data telah berhasil disimpan",
        html:
          "Total belanja : " + new Intl.NumberFormat().format(totalkeranjang),
        showConfirmButton: false,
        timer: 5000,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
      dispatch("handleInvoice", "call");
      reset();
    } else {
      css_loader = "display:none;";
      if (json.status == "500" || json.status == "404") {
        group_btn_beli = true;
        Swal.fire({
          position: "center",
          icon: "error",
          title: json.message,
          showConfirmButton: false,
          timer: 3000,
          background: daylight ? "#fff" : "#171717",
          color: daylight ? "#00a86b" : "#ff9900",
          toast: true,
        });
      } else {
        Swal.fire({
          icon: "error",
          title: "Oops...",
          text: json.message,
          background: daylight ? "#fff" : "#171717",
          color: daylight ? "#00a86b" : "#ff9900",
        });
      }
    }
  }

  function reverseString(str) {
    return str.split("").reverse().join("");
  }

  function reset() {
    keranjang = [];
    group_btn_beli = true;
    totalkeranjang = 0;
    count_line_5050umum = 0;
    count_line_5050special = 0;
    count_line_5050kombinasi = 0;
  }
  inittogel_432d("5050");
  function addKeranjang(
    nomor,
    game,
    bet,
    diskon_percen,
    diskon,
    bayar,
    win,
    kei_percen,
    kei,
    tipetoto
  ) {
    let total_data = keranjang.length;
    let flag_data = false;
    for (var i = 0; i < total_data; i++) {
      switch (game) {
        case "50_50_UMUM":
          if (nomor == keranjang[i].nomor.toString()) {
            let maxtotal_bayar5050umum = 0;
            for (var j = 0; j < keranjang.length; j++) {
              if ("50_50_UMUM" == keranjang[j].permainan) {
                if (parseInt(nomor) == parseInt(keranjang[j].nomor)) {
                  maxtotal_bayar5050umum =
                    parseInt(maxtotal_bayar5050umum) +
                    (parseInt(keranjang[j].bet) + parseInt(bet));
                }
              }
            }
            if (
              parseInt(limittotal_bet_5050umum) <
              parseInt(maxtotal_bayar5050umum)
            ) {
              temp_bulk_error +=
                "Nomor ini : " +
                nomor +
                " sudah melebihi LIMIT TOTAL 50-50 UMUM<br />";
              flag_data = true;
            }
          }
          break;
        case "50_50_SPECIAL":
          if (nomor == keranjang[i].nomor.toString()) {
            let maxtotal_bayar5050special = 0;
            for (var j = 0; j < keranjang.length; j++) {
              if ("50_50_UMUM" == keranjang[j].permainan) {
                if (parseInt(nomor) == parseInt(keranjang[j].nomor)) {
                  maxtotal_bayar5050special =
                    parseInt(maxtotal_bayar5050special) +
                    (parseInt(keranjang[j].bet) + parseInt(bet));
                }
              }
            }
            if (
              parseInt(limittotal_bet_5050special) <
              parseInt(maxtotal_bayar5050special)
            ) {
              temp_bulk_error +=
                "Nomor ini : " +
                nomor +
                " sudah melebihi LIMIT TOTAL 50-50 SPECIAL<br />";
              flag_data = true;
            }
          }
          break;
        case "50_50_KOMBINASI":
          if (nomor == keranjang[i].nomor.toString()) {
            let maxtotal_bayar5050kombinasi = 0;
            for (var j = 0; j < keranjang.length; j++) {
              if ("50_50_UMUM" == keranjang[j].permainan) {
                if (parseInt(nomor) == parseInt(keranjang[j].nomor)) {
                  maxtotal_bayar5050kombinasi =
                    parseInt(maxtotal_bayar5050kombinasi) +
                    (parseInt(keranjang[j].bet) + parseInt(bet));
                }
              }
            }
            if (
              parseInt(limittotal_bet_5050kombinasi) <
              parseInt(maxtotal_bayar5050kombinasi)
            ) {
              temp_bulk_error +=
                "Nomor ini : " +
                nomor +
                " sudah melebihi LIMIT TOTAL 50-50 KOMBINASI<br />";
              flag_data = true;
            }
          }
          break;
      }
    }

    if (flag_data == false) {
      nomor_global = nomor_global + 1;
      const data = {
        id: nomor_global,
        nomor,
        permainan: game,
        bet,
        diskon,
        diskonpercen: diskon_percen,
        bayar,
        win,
        kei,
        kei_percen,
        tipetoto,
      };
      keranjang = [data, ...keranjang];
      count_keranjang();
    }
  }
  const removekeranjang = (e) => {
    keranjang = keranjang.filter(
      (keranjang) => keranjang.id != e.detail.idkeranjang
    );
    totalkeranjang = totalkeranjang - e.detail.bayar;
    count_keranjang();
  };
  const removekeranjangall = (e) => {
    if (keranjang.length > 0) {
      reset();
      count_keranjang();
    } else {
      Swal.fire({
        position: "center",
        icon: "warning",
        title: "Tidak ada list transaksi",
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
  };
  const handleSave = (e) => {
    if (keranjang.length > 0) {
      swalWithBootstrapButtons
        .fire({
          title: "Apakah anda ingin melanjutkan?",
          html:
            "Total belanja anda sebesar : <strong>IDR. " +
            new Intl.NumberFormat().format(totalkeranjang) +
            "</strong>",
          icon: "warning",
          showCancelButton: true,
          confirmButtonText: "Ya, beli sekarang!",
          cancelButtonText: "Tidak!",
          reverseButtons: true,
          background: daylight ? "#fff" : "#171717",
          color: daylight ? "#00a86b" : "#ff9900",
        })
        .then((result) => {
          if (result.isConfirmed) {
            savetransaksi();
          } else if (
            /* Read more about handling dismissals below */
            result.dismiss === Swal.DismissReason.cancel
          ) {
            Swal.fire({
              position: "center",
              icon: "info",
              title: "Transaksi dibatalkan",
              showConfirmButton: false,
              timer: 3000,
              background: daylight ? "#fff" : "#171717",
              color: daylight ? "#00a86b" : "#ff9900",
              toast: true,
            });
          }
        });
    } else {
      Swal.fire({
        position: "center",
        icon: "warning",
        title: "Tidak ada list transaksi",
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
  };
  function count_keranjang() {
    let count_umum = 0;
    let count_special = 0;
    let count_kombinasi = 0;
    for (let i = 0; i < keranjang.length; i++) {
      switch (keranjang[i].permainan.toString()) {
        case "50_50_UMUM":
          count_umum = count_umum + 1;
          break;
        case "50_50_SPECIAL":
          count_special = count_special + 1;
          break;
        case "50_50_KOMBINASI":
          count_kombinasi = count_kombinasi + 1;
          break;
      }
    }
    count_line_5050umum = count_umum + db_form5050_umum_count_temp;
    count_line_5050special = count_special + db_form5050_special_count_temp;
    count_line_5050kombinasi =
      count_kombinasi + db_form5050_kombinasi_count_temp;
  }

  //5050 UMUM - INIT FORM
  let select_5050umum = "";
  let select_5050umum_input;
  let bet_5050umum = "";

  //5050 SPECIAL - INIT FORM
  let select_5050special_1 = "";
  let select_5050special_1_input;
  let select_5050special_2 = "";
  let select_5050special_2_input;
  let bet_5050special = "";

  //5050 KOMBINASI - INIT FORM
  let select_5050kombinasi_1 = "";
  let select_5050kombinasi_1_input;
  let select_5050kombinasi_2 = "";
  let select_5050kombinasi_2_input;
  let bet_5050kombinasi = "";

  function form_clear(e) {
    switch (e) {
      case "5050umum":
        select_5050umum = "";
        select_5050umum_input.focus();
        bet_5050umum = 0;
        break;
      case "5050special":
        select_5050special_1 = "";
        select_5050special_1_input.focus();
        select_5050special_2 = "";
        bet_5050special = 0;
        break;
      case "5050kombinasi":
        select_5050kombinasi_1 = "";
        select_5050kombinasi_1_input.focus();
        select_5050kombinasi_2 = "";
        bet_5050kombinasi = 0;
        break;
    }
  }
  function form5050_add() {
    let flag = true;
    let nomor = select_5050umum;
    let bet = bet_5050umum;
    let diskon = 0;
    let diskonpercen = 0;
    let kei = 0;
    let keipersen = 0;
    let bayar = 0;
    let win = 1;
    let nmgame = "50_50_UMUM";

    if (nomor == "") {
      select_5050umum_input.focus();
      flag = false;
    }

    if (bet == "") {
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Amount tidak boleh kosong",
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
    if (
      parseInt(bet) < parseInt(min_bet_5050umum) ||
      parseInt(bet) > parseInt(max_bet_5050umum)
    ) {
      bet_5050umum = min_bet_5050umum;
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title:
          "Minimal Bet : " +
          min_bet_5050umum +
          " Maximal Bet : " +
          max_bet_5050umum,
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }

    if (flag == true) {
      switch (nomor) {
        case "BESAR":
          kei = keibesar_bet_5050umum;
          diskon = discbesar_bet_5050umum;
          break;
        case "KECIL":
          kei = keikecil_bet_5050umum;
          diskon = disckecil_bet_5050umum;
          break;
        case "GENAP":
          kei = keigenap_bet_5050umum;
          diskon = discgenap_bet_5050umum;
          break;
        case "GANJIL":
          kei = keiganjil_bet_5050umum;
          diskon = discganjil_bet_5050umum;
          break;
        case "TENGAH":
          kei = keitengah_bet_5050umum;
          diskon = disctengah_bet_5050umum;
          break;
        case "TEPI":
          kei = keitepi_bet_5050umum;
          diskon = disctepi_bet_5050umum;
          break;
      }
      keipersen = kei;
      diskonpercen = diskon;
      kei = parseInt(bet) * kei;
      diskon = parseInt(bet) * diskon;
      bayar = parseInt(bet) - parseInt(diskon) - parseInt(kei);
      totalkeranjang = bayar + totalkeranjang;

      addKeranjang(
        nomor,
        nmgame,
        bet_5050umum,
        diskonpercen,
        diskon,
        bayar,
        win,
        keipersen,
        kei,
        flag_fulldiskon
      );
      form_clear("5050umum");
    }
  }
  function form5050special_add() {
    var flag = true;
    var nomor = select_5050special_1;
    var nomor2 = select_5050special_2;
    var bet = bet_5050special;
    var diskon = 0;
    var diskonpercen = 0;
    var win = 1;
    var bayar = 0;
    var kei = 0;
    var keipersen = 0;
    var nmgame = "50_50_SPECIAL";
    var msg = "";
    if (nomor == "") {
      select_5050umum_input.focus();
      flag = false;
    }
    if (nomor2 == "") {
      flag = false;
    }

    if (bet == "") {
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Amount tidak boleh kosong",
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
    if (parseInt(bet) < parseInt(min_bet_5050special)) {
      bet_5050special = min_bet_5050special;
      flag = false;
      msg += "Minimal Bet : " + min_bet_5050special;
    }
    if (parseInt(bet) > parseInt(max_bet_5050special)) {
      bet_5050special = max_bet_5050special;
      flag = false;
      msg += " Maximal Bet : " + max_bet_5050special;
    }
    if (flag == true) {
      switch (nomor + nomor2) {
        case "ASGANJIL":
          kei = keiasganjil_bet_5050special;
          diskon = discasganjil_bet_5050special;
          break;
        case "ASGENAP":
          kei = keiasgenap_bet_5050special;
          diskon = discasgenap_bet_5050special;
          break;
        case "ASBESAR":
          kei = keiasbesar_bet_5050special;
          diskon = discasbesar_bet_5050special;
          break;
        case "ASKECIL":
          kei = keiaskecil_bet_5050special;
          diskon = discaskecil_bet_5050special;
          break;
        case "KOPGANJIL":
          kei = keikopganjil_bet_5050special;
          diskon = disckopganjil_bet_5050special;
          break;
        case "KOPGENAP":
          kei = keikopgenap_bet_5050special;
          diskon = disckopgenap_bet_5050special;
          break;
        case "KOPBESAR":
          kei = keikopbesar_bet_5050special;
          diskon = disckopbesar_bet_5050special;
          break;
        case "KOPKECIL":
          kei = keikopkecil_bet_5050special;
          diskon = disckopkecil_bet_5050special;
          break;
        case "KEPALAGANJIL":
          kei = keikepalaganjil_bet_5050special;
          diskon = disckepalaganjil_bet_5050special;
          break;
        case "KEPALAGENAP":
          kei = keikepalagenap_bet_5050special;
          diskon = disckepalagenap_bet_5050special;
          break;
        case "KEPALABESAR":
          kei = keikepalabesar_bet_5050special;
          diskon = disckepalabesar_bet_5050special;
          break;
        case "KEPALAKECIL":
          kei = keikepalakecil_bet_5050special;
          diskon = disckepalakecil_bet_5050special;
          break;
        case "EKORGANJIL":
          kei = keiekorganjil_bet_5050special;
          diskon = discekorganjil_bet_5050special;
          break;
        case "EKORGENAP":
          kei = keiekorgenap_bet_5050special;
          diskon = discekorgenap_bet_5050special;
          break;
        case "EKORBESAR":
          kei = keiekorbesar_bet_5050special;
          diskon = discekorbesar_bet_5050special;
          break;
        case "EKORKECIL":
          kei = keiekorkecil_bet_5050special;
          diskon = discekorkecil_bet_5050special;
          break;
      }
      keipersen = kei;
      diskonpercen = diskon;
      kei = parseInt(bet) * kei;
      diskon = parseInt(bet) * diskon;
      bayar = parseInt(bet) - parseInt(diskon) - parseInt(kei);
      totalkeranjang = bayar + totalkeranjang;

      addKeranjang(
        nomor + "_" + nomor2,
        nmgame,
        bet_5050special,
        diskonpercen,
        diskon,
        bayar,
        win,
        keipersen,
        kei,
        flag_fulldiskon
      );
      form_clear("5050special");
    }

    if (msg != "") {
      Swal.fire({
        position: "center",
        icon: "warning",
        title: msg,
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
  }
  function form5050kombinasi_add() {
    var flag = true;
    var nomor = select_5050kombinasi_1;
    var nomor2 = select_5050kombinasi_2;
    var bet = bet_5050kombinasi;
    var diskon = 0;
    var diskonpercen = 0;
    var win = 1;
    var bayar = 0;
    var kei = 0;
    var keipersen = 0;
    var nmgame = "50_50_KOMBINASI";

    if (nomor == "") {
      select_5050special_1_input.focus();
      flag = false;
    }
    if (nomor2 == "") {
      flag = false;
    }

    if (bet == "") {
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Amount tidak boleh kosong",
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
    if (parseInt(bet) < parseInt(min_bet_5050kombinasi)) {
      bet_5050kombinasi = min_bet_5050kombinasi;
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Minimal Bet : " + min_bet_5050kombinasi,
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
    if (parseInt(bet) > parseInt(max_bet_5050kombinasi)) {
      bet_5050kombinasi = max_bet_5050kombinasi;
      flag = false;
      Swal.fire({
        position: "center",
        icon: "error",
        title: " Maximal Bet : " + max_bet_5050kombinasi,
        showConfirmButton: false,
        timer: 1500,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
    if (flag == true) {
      switch (nomor + nomor2) {
        case "BELAKANGMONO":
          kei = kei_belakangmono_bet_5050kombinasi;
          diskon = disc_belakangmono_bet_5050kombinasi;
          break;
        case "BELAKANGSTEREO":
          kei = kei_belakangstereo_bet_5050kombinasi;
          diskon = disc_belakangstereo_bet_5050kombinasi;
          break;
        case "BELAKANGKEMBANG":
          kei = kei_belakangkembang_bet_5050kombinasi;
          diskon = disc_belakangkembang_bet_5050kombinasi;
          break;
        case "BELAKANGKEMPIS":
          kei = kei_belakangkempis_bet_5050kombinasi;
          diskon = disc_belakangkempis_bet_5050kombinasi;
          break;
        case "BELAKANGKEMBAR":
          kei = kei_belakangkembar_bet_5050kombinasi;
          diskon = disc_belakangkembar_bet_5050kombinasi;
          break;
        case "TENGAHMONO":
          kei = kei_tengahmono_bet_5050kombinasi;
          diskon = disc_tengahmono_bet_5050kombinasi;
          break;
        case "TENGAHSTEREO":
          kei = kei_tengahstereo_bet_5050kombinasi;
          diskon = disc_tengahstereo_bet_5050kombinasi;
          break;
        case "TENGAHKEMBANG":
          kei = kei_tengahkembang_bet_5050kombinasi;
          diskon = disc_tengahkembang_bet_5050kombinasi;
          break;
        case "TENGAHKEMPIS":
          kei = kei_tengahkempis_bet_5050kombinasi;
          diskon = disc_tengahkempis_bet_5050kombinasi;
          break;
        case "TENGAHKEMBAR":
          kei = kei_tengahkembar_bet_5050kombinasi;
          diskon = disc_tengahkembar_bet_5050kombinasi;
          break;
        case "DEPANMONO":
          kei = kei_depanmono_bet_5050kombinasi;
          diskon = disc_depanmono_bet_5050kombinasi;
          break;
        case "DEPANSTEREO":
          kei = kei_depanstereo_bet_5050kombinasi;
          diskon = disc_depanstereo_bet_5050kombinasi;
          break;
        case "DEPANKEMBANG":
          kei = kei_depankembang_bet_5050kombinasi;
          diskon = disc_depankembang_bet_5050kombinasi;
          break;
        case "DEPANKEMPIS":
          kei = kei_depankempis_bet_5050kombinasi;
          diskon = disc_depankempis_bet_5050kombinasi;
          break;
        case "DEPANKEMBAR":
          kei = kei_depankembar_bet_5050kombinasi;
          diskon = disc_depankembar_bet_5050kombinasi;
          break;
      }
      keipersen = kei;
      diskonpercen = diskon;
      kei = parseInt(bet) * kei;
      diskon = parseInt(bet) * diskon;
      bayar = parseInt(bet) - parseInt(kei);

      totalkeranjang = bayar + totalkeranjang;

      addKeranjang(
        nomor + "_" + nomor2,
        nmgame,
        bet_5050kombinasi,
        diskonpercen,
        diskon,
        bayar,
        win,
        keipersen,
        kei,
        flag_fulldiskon
      );
      form_clear("5050kombinasi");
    }
  }

  const handleTambah = (e) => {
    switch (e) {
      case "5050umum":
        if (
          select_5050umum == "" &&
          parseInt(bet_5050umum) < min_bet_5050umum
        ) {
          select_5050special_1_input.focus();
        } else {
          form5050_add();
        }
        break;
      case "5050special":
        if (
          select_5050special_1 == "" &&
          select_5050special_2 == "" &&
          parseInt(bet_5050umum) < min_bet_5050special
        ) {
          select_5050special_1_input.focus();
        } else {
          form5050special_add();
        }
        break;
      case "5050kombinasi":
        if (
          select_5050kombinasi_1 == "" &&
          select_5050kombinasi_2 == "" &&
          parseInt(bet_5050kombinasi) < min_bet_5050kombinasi
        ) {
          select_5050kombinasi_1_input.focus();
        } else {
          form5050kombinasi_add();
        }
        break;
    }
  };

  const handleKeyboard_number = (e) => {
    let numbera;
    for (let i = 0; i < bet_5050umum.length; i++) {
      numbera = parseInt(bet_5050umum[i]);
      if (isNaN(numbera)) {
        bet_5050umum = "";
      }
    }
    for (let i = 0; i < bet_5050special.length; i++) {
      numbera = parseInt(bet_5050special[i]);
      if (isNaN(numbera)) {
        bet_5050special = "";
      }
    }
    for (let i = 0; i < bet_5050kombinasi.length; i++) {
      numbera = parseInt(bet_5050kombinasi[i]);
      if (isNaN(numbera)) {
        bet_5050kombinasi = "";
      }
    }
  };
  const handleKeyboard_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      form5050_add();
    }
  };
  const handleKeyboardspecial_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      form5050special_add();
    }
  };
  const handleKeyboardkombinasi_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      form5050kombinasi_add();
    }
  };
</script>

<Loader cssstyle={css_loader} />
<Card class={daylight ? "" : "bg-dark"} style="margin:0px;padding:0px;">
  <PeriodePanel
    {pasaran_name}
    {permainan_title}
    {pasaran_periode}
    {pasaran_code}
    {daylight}
    {client_device}
  />
  <CardBody class={daylight ? "" : "dark"}>
    <TabContent class="periode-menu {daylight ? '' : 'dark'}">
      <!-- Umum -->
      <TabPane tabId="form_5050umum" tab="UMUM" active>
        <div style="margin:10px 0;">
          <div class="row gap-3">
            <div class="col-md">
              <div class="form-floating">
                <select
                  class="form-select button-bet-default"
                  class:dark={daylight === false}
                  bind:value={select_5050umum}
                  bind:this={select_5050umum_input}
                  id="selectOpt1"
                  aria-label="Floating label select"
                >
                  <option selected>--Pilih--</option>
                  <option value="BESAR">BESAR</option>
                  <option value="KECIL">KECIL</option>
                  <option value="GENAP">GENAP</option>
                  <option value="GANJIL">GANJIL</option>
                  <option value="TENGAH">TENGAH</option>
                  <option value="TEPI">TEPI</option>
                </select>
                <label for="selectOpt1">Pilih</label>
              </div>
            </div>
            <div class="col-md">
              <div class="form-floating">
                <input
                  bind:value={bet_5050umum}
                  on:keyup={handleKeyboard_number}
                  on:keypress={handleKeyboard_checkenter}
                  type="text"
                  id="betQuick2D"
                  class="form-control fs-5 text-end button-bet-default"
                  class:dark={daylight === false}
                  placeholder="Bet"
                  minlength="3"
                  maxlength="7"
                  tab_index="0"
                  autocomplete="off"
                />
                <span style="float:right;font-size:12px;color:#8a8a8a;"
                  >{new Intl.NumberFormat().format(bet_5050umum)}</span
                >
                <label for="betQuick2D" class="form-label"
                  >Bet (min : {new Intl.NumberFormat().format(min_bet_5050umum)}
                  dan max : {new Intl.NumberFormat().format(
                    max_bet_5050umum
                  )})</label
                >
              </div>
            </div>
            <div class="col-md">
              <Button
                id="btn2"
                class="form-control mt-2"
                style="border-radius:5px"
                on:click={() => {
                  handleTambah("5050umum");
                }}>TAMBAH</Button
              >
            </div>
          </div>
        </div>
      </TabPane>

      <!-- Special -->
      <TabPane tabId="form_5050special" tab="SPECIAL">
        <div style="margin:10px 0;">
          <div class="row gap-3">
            <div class="col-md">
              <div class="form-floating">
                <select
                  class="form-select button-bet-default"
                  class:dark={daylight === false}
                  bind:value={select_5050special_1}
                  bind:this={select_5050special_1_input}
                  id="selectSpc1"
                  aria-label="Floating label select"
                  required
                >
                  <option selected>--Pilih--</option>
                  <option value="AS">AS</option>
                  <option value="KOP">KOP</option>
                  <option value="KEPALA">KEPALA</option>
                  <option value="EKOR">EKOR</option>
                </select>
                <label for="selectSpc1">Pilih</label>
              </div>
            </div>
            <div class="col-md">
              <div class="form-floating">
                <select
                  class="form-select button-bet-default"
                  class:dark={daylight === false}
                  bind:value={select_5050special_2}
                  bind:this={select_5050special_2_input}
                  id="selectSpc2"
                  aria-label="Floating label select"
                  required
                >
                  <option selected>--Pilih--</option>
                  <option value="GENAP">GENAP</option>
                  <option value="GANJIL">GANJIL</option>
                  <option value="BESAR">BESAR</option>
                  <option value="KECIL">KECIL</option>
                </select>
                <label for="selectSpc2">Pilih</label>
              </div>
            </div>
            <div class="col-md">
              <div class="form-floating">
                <input
                  bind:value={bet_5050special}
                  on:keyup={handleKeyboard_number}
                  on:keypress={handleKeyboardspecial_checkenter}
                  type="text"
                  id="inptSpc"
                  class="form-control fs-5 text-end button-bet-default"
                  class:dark={daylight === false}
                  placeholder="Bet"
                  minlength="3"
                  maxlength="7"
                  tab_index="0"
                  autocomplete="off"
                />
                <span style="text-align:right;font-size:12px;color:#8a8a8a;"
                  >{new Intl.NumberFormat().format(bet_5050special)}</span
                >
                <label for="inptSpc" class="form-label"
                  >Bet (min : {new Intl.NumberFormat().format(
                    min_bet_5050special
                  )} dan max : {new Intl.NumberFormat().format(
                    max_bet_5050special
                  )})</label
                >
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col">
              <Button
                id="btn2"
                class="form-control mt-2"
                style="border-radius:5px"
                on:click={() => {
                  handleTambah("5050special");
                }}>TAMBAH</Button
              >
            </div>
          </div>
        </div>
      </TabPane>

      <!-- Kombinasi -->
      <TabPane tabId="form_5050kombinasi" tab="KOMBINASI">
        <div style="margin:10px 0;">
          <div class="row gap-3">
            <div class="col-md">
              <div class="form-floating">
                <select
                  class="form-select button-bet-default"
                  class:dark={daylight === false}
                  bind:value={select_5050kombinasi_1}
                  bind:this={select_5050kombinasi_1_input}
                  id="selectKmbn1"
                  aria-label="Floating label select"
                >
                  <option value="">--Pilih--</option>
                  <option value="BELAKANG">BELAKANG</option>
                  <option value="TENGAH">TENGAH</option>
                  <option value="DEPAN">DEPAN</option>
                </select>
                <label for="selectKmbn1">Pilih</label>
              </div>
            </div>
            <div class="col-md">
              <div class="form-floating">
                <select
                  class="form-select button-bet-default"
                  class:dark={daylight === false}
                  bind:value={select_5050kombinasi_2}
                  bind:this={select_5050kombinasi_2_input}
                  id="selectKmbn2"
                  aria-label="Floating label select"
                >
                  <option value="">--Pilih--</option>
                  <option value="MONO">MONO</option>
                  <option value="STEREO">STEREO</option>
                  <option value="KEMBANG">KEMBANG</option>
                  <option value="KEMPIS">KEMPIS</option>
                  <option value="KEMBAR">KEMBAR</option>
                </select>
                <label for="selectKmbn2">Pilih</label>
              </div>
            </div>
            <div class="col-md">
              <div class="form-floating">
                <input
                  bind:value={bet_5050kombinasi}
                  on:keyup={handleKeyboard_number}
                  on:keypress={handleKeyboardkombinasi_checkenter}
                  type="text"
                  id="inptKmbn"
                  class="form-control fs-5 text-end button-bet-default"
                  class:dark={daylight === false}
                  placeholder="Bet"
                  minlength="3"
                  maxlength="7"
                  tab_index="0"
                  autocomplete="off"
                />
                <span style="text-align:right;font-size:12px;color:#8a8a8a;"
                  >{new Intl.NumberFormat().format(bet_5050kombinasi)}</span
                >
                <label for="inptKmbn" class="form-label text-truncate"
                  >Bet (min : {new Intl.NumberFormat().format(
                    min_bet_5050special
                  )} dan max : {new Intl.NumberFormat().format(
                    max_bet_5050special
                  )})</label
                >
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col">
              <Button
                id="btn2"
                class="form-control mt-2"
                style="border-radius:5px"
                on:click={() => {
                  handleTambah("5050kombinasi");
                }}>TAMBAH</Button
              >
            </div>
          </div>
        </div>
      </TabPane>
    </TabContent>
  </CardBody>
</Card>

<Modal modal_id={"modalError"} modal_size={"modal-dialog-centered"}>
  <slot:template slot="header">
    <div class="float-end">
      <img
        style="cursor:pointer;"
        data-bs-dismiss="modal"
        aria-label="Close"
        width="25"
        src="https://i.ibb.co/9yNF25L/outline-close-white-48dp.png"
        alt=""
      />
    </div>
    <h5 class="modal-title" id="exampleModalLabel">ERROR</h5>
  </slot:template>
  <slot:template slot="body">
    {@html temp_bulk_error}
  </slot:template>
</Modal>
<div class="clearfix" />
<br />
<Tablekeranjang5050
  on:removekeranjang={removekeranjang}
  on:removekeranjangall={removekeranjangall}
  on:handleSave={handleSave}
  {client_device}
  {group_btn_beli}
  {count_line_5050umum}
  {count_line_5050special}
  {count_line_5050kombinasi}
  {keranjang}
  {totalkeranjang}
  {min_bet_5050umum}
  {max_bet_5050umum}
  {keibesar_bet_5050umum}
  {keikecil_bet_5050umum}
  {keigenap_bet_5050umum}
  {keiganjil_bet_5050umum}
  {keitengah_bet_5050umum}
  {keitepi_bet_5050umum}
  {discbesar_bet_5050umum}
  {disckecil_bet_5050umum}
  {discgenap_bet_5050umum}
  {discganjil_bet_5050umum}
  {disctengah_bet_5050umum}
  {disctepi_bet_5050umum}
  {min_bet_5050special}
  {max_bet_5050special}
  {keiasganjil_bet_5050special}
  {keiasgenap_bet_5050special}
  {keiasbesar_bet_5050special}
  {keiaskecil_bet_5050special}
  {keikopganjil_bet_5050special}
  {keikopgenap_bet_5050special}
  {keikopbesar_bet_5050special}
  {keikopkecil_bet_5050special}
  {keikepalaganjil_bet_5050special}
  {keikepalagenap_bet_5050special}
  {keikepalabesar_bet_5050special}
  {keikepalakecil_bet_5050special}
  {keiekorganjil_bet_5050special}
  {keiekorgenap_bet_5050special}
  {keiekorbesar_bet_5050special}
  {keiekorkecil_bet_5050special}
  {discasganjil_bet_5050special}
  {discasgenap_bet_5050special}
  {discasbesar_bet_5050special}
  {discaskecil_bet_5050special}
  {disckopganjil_bet_5050special}
  {disckopgenap_bet_5050special}
  {disckopbesar_bet_5050special}
  {disckopkecil_bet_5050special}
  {disckepalaganjil_bet_5050special}
  {disckepalagenap_bet_5050special}
  {disckepalabesar_bet_5050special}
  {disckepalakecil_bet_5050special}
  {discekorganjil_bet_5050special}
  {discekorgenap_bet_5050special}
  {discekorbesar_bet_5050special}
  {discekorkecil_bet_5050special}
  {min_bet_5050kombinasi}
  {max_bet_5050kombinasi}
  {kei_belakangmono_bet_5050kombinasi}
  {kei_belakangstereo_bet_5050kombinasi}
  {kei_belakangkembang_bet_5050kombinasi}
  {kei_belakangkempis_bet_5050kombinasi}
  {kei_belakangkembar_bet_5050kombinasi}
  {kei_tengahmono_bet_5050kombinasi}
  {kei_tengahstereo_bet_5050kombinasi}
  {kei_tengahkembang_bet_5050kombinasi}
  {kei_tengahkempis_bet_5050kombinasi}
  {kei_tengahkembar_bet_5050kombinasi}
  {kei_depanmono_bet_5050kombinasi}
  {kei_depanstereo_bet_5050kombinasi}
  {kei_depankembang_bet_5050kombinasi}
  {kei_depankempis_bet_5050kombinasi}
  {kei_depankembar_bet_5050kombinasi}
  {disc_belakangmono_bet_5050kombinasi}
  {disc_belakangstereo_bet_5050kombinasi}
  {disc_belakangkembang_bet_5050kombinasi}
  {disc_belakangkempis_bet_5050kombinasi}
  {disc_belakangkembar_bet_5050kombinasi}
  {disc_tengahmono_bet_5050kombinasi}
  {disc_tengahstereo_bet_5050kombinasi}
  {disc_tengahkembang_bet_5050kombinasi}
  {disc_tengahkempis_bet_5050kombinasi}
  {disc_tengahkembar_bet_5050kombinasi}
  {disc_depanmono_bet_5050kombinasi}
  {disc_depanstereo_bet_5050kombinasi}
  {disc_depankembang_bet_5050kombinasi}
  {disc_depankempis_bet_5050kombinasi}
  {disc_depankembar_bet_5050kombinasi}
  {daylight}
/>

<style>
  input,
  input::-webkit-input-placeholder {
    font-size: 13px;
    color: #8a8a8a;
  }
  input::placeholder {
    color: #8a8a8a;
  }
</style>

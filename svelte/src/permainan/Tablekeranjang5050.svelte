<script>
  import { Button, Icon, TabContent, TabPane } from "sveltestrap";
  import { createEventDispatcher } from "svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Modal from "../components/Modalfull2.svelte";
  import Lineinfo from "../components/Keranjang.svelte";

  export let keranjang = [];
  export let client_device = "";
  export let totalkeranjang = 0;
  export let group_btn_beli = false;
  export let count_line_5050umum = 0;
  export let count_line_5050special = 0;
  export let count_line_5050kombinasi = 0;
  export let min_bet_5050umum = 0;
  export let max_bet_5050umum = 0;
  export let keibesar_bet_5050umum = 0;
  export let keikecil_bet_5050umum = 0;
  export let keigenap_bet_5050umum = 0;
  export let keiganjil_bet_5050umum = 0;
  export let keitengah_bet_5050umum = 0;
  export let keitepi_bet_5050umum = 0;
  export let discbesar_bet_5050umum = 0;
  export let disckecil_bet_5050umum = 0;
  export let discgenap_bet_5050umum = 0;
  export let discganjil_bet_5050umum = 0;
  export let disctengah_bet_5050umum = 0;
  export let disctepi_bet_5050umum = 0;
  export let min_bet_5050special = 0;
  export let max_bet_5050special = 0;
  export let keiasganjil_bet_5050special = 0;
  export let keiasgenap_bet_5050special = 0;
  export let keiasbesar_bet_5050special = 0;
  export let keiaskecil_bet_5050special = 0;
  export let keikopganjil_bet_5050special = 0;
  export let keikopgenap_bet_5050special = 0;
  export let keikopbesar_bet_5050special = 0;
  export let keikopkecil_bet_5050special = 0;
  export let keikepalaganjil_bet_5050special = 0;
  export let keikepalagenap_bet_5050special = 0;
  export let keikepalabesar_bet_5050special = 0;
  export let keikepalakecil_bet_5050special = 0;
  export let keiekorganjil_bet_5050special = 0;
  export let keiekorgenap_bet_5050special = 0;
  export let keiekorbesar_bet_5050special = 0;
  export let keiekorkecil_bet_5050special = 0;
  export let discasganjil_bet_5050special = 0;
  export let discasgenap_bet_5050special = 0;
  export let discasbesar_bet_5050special = 0;
  export let discaskecil_bet_5050special = 0;
  export let disckopganjil_bet_5050special = 0;
  export let disckopgenap_bet_5050special = 0;
  export let disckopbesar_bet_5050special = 0;
  export let disckopkecil_bet_5050special = 0;
  export let disckepalaganjil_bet_5050special = 0;
  export let disckepalagenap_bet_5050special = 0;
  export let disckepalabesar_bet_5050special = 0;
  export let disckepalakecil_bet_5050special = 0;
  export let discekorganjil_bet_5050special = 0;
  export let discekorgenap_bet_5050special = 0;
  export let discekorbesar_bet_5050special = 0;
  export let discekorkecil_bet_5050special = 0;
  export let min_bet_5050kombinasi = 0;
  export let max_bet_5050kombinasi = 0;
  export let kei_belakangmono_bet_5050kombinasi = 0;
  export let kei_belakangstereo_bet_5050kombinasi = 0;
  export let kei_belakangkembang_bet_5050kombinasi = 0;
  export let kei_belakangkempis_bet_5050kombinasi = 0;
  export let kei_belakangkembar_bet_5050kombinasi = 0;
  export let kei_tengahmono_bet_5050kombinasi = 0;
  export let kei_tengahstereo_bet_5050kombinasi = 0;
  export let kei_tengahkembang_bet_5050kombinasi = 0;
  export let kei_tengahkempis_bet_5050kombinasi = 0;
  export let kei_tengahkembar_bet_5050kombinasi = 0;
  export let kei_depanmono_bet_5050kombinasi = 0;
  export let kei_depanstereo_bet_5050kombinasi = 0;
  export let kei_depankembang_bet_5050kombinasi = 0;
  export let kei_depankempis_bet_5050kombinasi = 0;
  export let kei_depankembar_bet_5050kombinasi = 0;
  export let disc_belakangmono_bet_5050kombinasi = 0;
  export let disc_belakangstereo_bet_5050kombinasi = 0;
  export let disc_belakangkembang_bet_5050kombinasi = 0;
  export let disc_belakangkempis_bet_5050kombinasi = 0;
  export let disc_belakangkembar_bet_5050kombinasi = 0;
  export let disc_tengahmono_bet_5050kombinasi = 0;
  export let disc_tengahstereo_bet_5050kombinasi = 0;
  export let disc_tengahkembang_bet_5050kombinasi = 0;
  export let disc_tengahkempis_bet_5050kombinasi = 0;
  export let disc_tengahkembar_bet_5050kombinasi = 0;
  export let disc_depanmono_bet_5050kombinasi = 0;
  export let disc_depanstereo_bet_5050kombinasi = 0;
  export let disc_depankembang_bet_5050kombinasi = 0;
  export let disc_depankempis_bet_5050kombinasi = 0;
  export let disc_depankembar_bet_5050kombinasi = 0;
  export let daylight;
  let count_line = 0;
  let dispatch = createEventDispatcher();

  const handleRemoveKeranjang = (idkeranjang, bayar) => {
    dispatch("removekeranjang", {
      idkeranjang,
      bayar,
    });
  };
  const handleRemoveKeranjang_all = () => {
    dispatch("removekeranjangall", "all");
  };
  const handleSave = () => {
    dispatch("handleSave", "save");
  };
  let open_informasi = false;
  let fullscreen = "";
  const toggle = () => {
    fullscreen = "xl";
    open_informasi = !open_informasi;
  };
  $: count_line =
    count_line_5050umum + count_line_5050special + count_line_5050kombinasi;
</script>

<PanelFull
  header={true}
  footer={true}
  body_style="padding:0px;background:{daylight
    ? '#fff'
    : '#181717'};height:450px;"
  {daylight}
>
  <slot:template slot="header">
    <div class="container px-0">
      <div class="row">
        <div class="col-md">
          <h1 style="padding:5px;margin:0px;font-size:15px;">
            TOTAL BAYAR : <span style="color:{daylight ? '#00A86B' : '#fc0'};"
              >{new Intl.NumberFormat().format(totalkeranjang)}</span
            >
          </h1>
        </div>

        <div class="col-md">
          <div class="float-end" id="btnbelitogel">
            <Button
              id="btn1"
              data-bs-toggle="modal"
              data-bs-target="#modalInformasi"
              style="margin-top:5px;"
              ><Icon name="info-circle" /> Informasi</Button
            >
            {#if group_btn_beli == true}
              <Button
                id="btn1"
                on:click={handleRemoveKeranjang_all}
                style="margin-top:5px;"
                ><Icon name="trash" /> Hapus Semua</Button
              >
              <Button id="btn2" on:click={handleSave} style="margin-top:5px;"
                ><Icon name="cart-check" /> BELI</Button
              >
            {/if}
          </div>
        </div>
      </div>
    </div>
  </slot:template>
  <slot:template slot="body">
    <table
      class="table {daylight
        ? ''
        : 'table-dark table-striped '} table-sm table-responsive"
    >
      <thead>
        <tr>
          <th
            width="1%"
            style="text-align:center;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>&nbsp;</th
          >
          <th
            width="10%"
            style="text-align:center;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>NOMOR</th
          >
          <th
            width="20%"
            style="text-align:center;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>PERMAINAN</th
          >
          <th
            width="20%"
            style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>BET</th
          >
          <th
            width="20%"
            style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>DISKON</th
          >
          <th
            width="20%"
            style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>KEI</th
          >
          <th
            width="*"
            style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
            NOWRAP>BAYAR</th
          >
        </tr>
      </thead>
      <tbody>
        {#each keranjang as rec}
          <tr>
            <td
              style="text-align:center;vertical-align:middle;"
              on:click={() => {
                handleRemoveKeranjang(rec.id, rec.bayar);
              }}
            >
              <Icon name="trash" style="cursor:pointer;" />
            </td>
            <td class="table-fill text-center" class:dark={daylight === false}
              >{rec.nomor}</td
            >
            <td class="table-fill text-center" class:dark={daylight === false}
              >{rec.permainan}</td
            >
            <td class="table-fill text-end" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(rec.bet)}
            </td>
            <td class="table-fill text-end" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(Math.ceil(rec.diskon))} ({(
                rec.diskonpercen * 100
              ).toFixed(1)}%)
            </td>
            <td class="table-fill text-end" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(Math.ceil(rec.kei))}
              ({(rec.kei_percen * 100).toFixed(1)}%)
            </td>
            <td class="table-fill text-end" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(rec.bayar)}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Lineinfo {daylight} {count_line} />
  </slot:template>
</PanelFull>

<Modal
  modal_id={"modalInformasi"}
  modal_footer_flag={false}
  modal_body_height={"height:500px;"}
  modal_size={"modal-dialog-centered"}
  {daylight}
>
  <slot:template slot="header">
    <h5 class="modal-title">50-50</h5>
  </slot:template>
  <slot:template slot="body">
    <TabContent
      class="periode-menu {daylight ? '' : 'dark'}"
      style="padding: 0px;margin:0px;"
    >
      <TabPane
        tabId="informasi_5050umum"
        tab="UMUM"
        active
        style="padding:0px;font-size:12px;color:#8a8a8a;"
      >
        <table class="table table-dark table-sm">
          <tbody>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MIN BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(min_bet_5050umum)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_5050umum)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                BESAR : {Number(discbesar_bet_5050umum * 100).toFixed(2)}%
                <br />
                KECIL : {Number(disckecil_bet_5050umum * 100).toFixed(2)}%
                <br />
                GENAP : {Number(discgenap_bet_5050umum * 100).toFixed(2)}%
                <br />
                GANJIL : {Number(discganjil_bet_5050umum * 100).toFixed(2)}%
                <br />
                TENGAH : {Number(disctengah_bet_5050umum * 100).toFixed(2)}%
                <br />
                TEPI : {Number(disctepi_bet_5050umum * 100).toFixed(2)}%
                <br />
              </td>
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >KEI</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                BESAR : {Number(keibesar_bet_5050umum * 100).toFixed(2)}% <br />
                KECIL : {Number(keikecil_bet_5050umum * 100).toFixed(2)}% <br />
                GENAP : {Number(keigenap_bet_5050umum * 100).toFixed(2)}% <br />
                GANJIL : {Number(keiganjil_bet_5050umum * 100).toFixed(2)}%
                <br />
                TENGAH : {Number(keitengah_bet_5050umum * 100).toFixed(2)}%
                <br />
                TEPI : {Number(keitepi_bet_5050umum * 100).toFixed(2)}%
                <br />
              </td>
            </tr>
          </tbody>
        </table>
        <p
          style="padding: 5px;color:{daylight
            ? '#171717'
            : 'white'};font-size:13px;"
        >
          CARA BERMAIN
          <br />
          Permainan ganjil/genap, besar/kecil, dan tengah/ tepi UMUM
          <br />
          Struktur: CD ( hanya dua angka belakang yang dihitung)<br /><br />

          Menebak ganjil/genap dan besar/kecil dari posisi:<br />
          C=KEPALA<br />
          D=EKOR<br />
          Besar/Kecil: 0-4=kecil, 5-9=besar<br />
          Ganjil/Genap : 1=ganjil, 2=genap dan seterusnya<br /><br />
          Tengah/Tepi : Tengah: angka 25 s/d 74. TEPI : angka 75 s/d 99, dan 00s/d
          24<br /><br />

          Analisis I :<br />
          Keluar : 6789, Berarti hasil 89<br />
          Berarti pemenang adalah yang memilih :<br />
          BESAR<br />
          GANJIL<br />
          TEPI
        </p>
      </TabPane>
      <TabPane
        tabId="informasi_5050special"
        tab="SPECIAL"
        style="padding:0px;font-size:12px;color:#8a8a8a;"
      >
        <table class="table table-dark table-sm">
          <tbody>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MIN BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(min_bet_5050special)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_5050special)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                AS GANJIL : {Number(discasganjil_bet_5050special * 100).toFixed(
                  2
                )}% <br />
                AS GENAP : {Number(discasgenap_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                AS BESAR : {Number(discasbesar_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                AS KECIL : {Number(discaskecil_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KOP GANIL : {Number(
                  disckopganjil_bet_5050special * 100
                ).toFixed(2)}% <br />
                KOP GENAP : {Number(disckopgenap_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KOP BESAR : {Number(disckopbesar_bet_5050special * 100).toFixed(
                  2
                )}% <br />
                KOP KECIL : {Number(disckopkecil_bet_5050special * 100).toFixed(
                  2
                )}% <br />
                KEPALA GANIL : {Number(
                  disckepalaganjil_bet_5050special * 100
                ).toFixed(2)}% <br />
                KEPALA GENAP : {Number(
                  disckepalagenap_bet_5050special * 100
                ).toFixed(2)}%
                <br />
                KEPALA BESAR : {Number(
                  disckepalabesar_bet_5050special * 100
                ).toFixed(2)}% <br />
                KEPALA KECIL : {Number(
                  disckepalakecil_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR GANIL : {Number(
                  discekorganjil_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR GENAP : {Number(
                  discekorgenap_bet_5050special * 100
                ).toFixed(2)}%
                <br />
                EKOR BESAR : {Number(
                  discekorbesar_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR KECIL : {Number(
                  discekorkecil_bet_5050special * 100
                ).toFixed(2)}% <br />
              </td>
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >KEI</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                AS GANJIL : {Number(keiasganjil_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                AS GENAP : {Number(keiasgenap_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                AS BESAR : {Number(keiasbesar_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                AS KECIL : {Number(keiaskecil_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KOP GANIL : {Number(keikopganjil_bet_5050special * 100).toFixed(
                  2
                )}% <br />
                KOP GENAP : {Number(keikopgenap_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KOP BESAR : {Number(keikopbesar_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KOP KECIL : {Number(keikopkecil_bet_5050special * 100).toFixed(
                  2
                )}%
                <br />
                KEPALA GANIL : {Number(
                  keikepalaganjil_bet_5050special * 100
                ).toFixed(2)}% <br />
                KEPALA GENAP : {Number(
                  keikepalagenap_bet_5050special * 100
                ).toFixed(2)}%
                <br />
                KEPALA BESAR : {Number(
                  keikepalabesar_bet_5050special * 100
                ).toFixed(2)}% <br />
                KEPALA KECIL : {Number(
                  keikepalakecil_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR GANIL : {Number(
                  keiekorganjil_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR GENAP : {Number(
                  keiekorgenap_bet_5050special * 100
                ).toFixed(2)}%
                <br />
                EKOR BESAR : {Number(
                  keiekorbesar_bet_5050special * 100
                ).toFixed(2)}% <br />
                EKOR KECIL : {Number(
                  keiekorkecil_bet_5050special * 100
                ).toFixed(2)}% <br />
              </td>
            </tr>
          </tbody>
        </table>
        <p
          style="padding: 5px;color:{daylight
            ? '#171717'
            : 'white'};font-size:13px;"
        >
          CARA BERMAIN
          <br />
          Menebak ganjil/genap, besar/kecil <br />
          Permainan ini sangat menarik karena pasarannya naik turun sesuai keinginan
          market pada waktu tersebut. Dengan demikian, nilai pembelian dipengaruhi
          kei (pasaran)
          <br />
          Struktur: ABCD<br /><br />

          Menebak ganjil/genap dan besar/kecil dari posisi:<br />
          A=AS<br />
          B=KOP<br />
          C=KEPALA<br />
          D=EKOR<br />

          Besar/Kecil : 0-4=kecil, 5-9=besar<br />
          Ganjil/Genap : 1=ganjil, 2=genap dan seterusnya<br /><br />

          Analisis I :<br />
          Keluar : 4327<br />
          Berarti pemenang adalah yang memilih :<br />
          AS GENAP/KECIL<br />
          KOP GANJIL/KECIL<br />
          KEPALA GENAP/KECIL<br />
          EKOR GANJIL/BESAR<br />
          Misal anda membeli dengan dana Rp.100rb untuk AS Genap, menang = 100rb
          + [indeks pasaran AS Genap 50-50]<br />
          Atau :<br />
          Jika membeli dengan dana Rp.100rb untuk Ekor Ganjil, menang = 100rb + [indeks
          pasaran Ekor Ganjil 50-50]<br />
          Atau :<br />
          Jika membeli dengan dana Rp.100rb untuk AS Ganjil, kalah = 100rb + [indeks
          pasaran AS Ganjil 50-50]
        </p>
      </TabPane>
      <TabPane
        tabId="informasi_5050kombinasi"
        tab="KOMBINASI"
        style="padding:0px;font-size:12px;color:#8a8a8a;"
      >
        <table class="table table-dark table-sm">
          <tbody>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MIN BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(min_bet_5050kombinasi)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_5050kombinasi)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                BELAKANG MONO : {Number(
                  disc_belakangmono_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG STEREO : {Number(
                  disc_belakangstereo_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG KEMBANG : {Number(
                  disc_belakangkembang_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
                BELAKANG KEMPIS : {Number(
                  disc_belakangkempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG KEMBAR : {Number(
                  disc_belakangkembar_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH MONO : {Number(
                  disc_tengahmono_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH STEREO : {Number(
                  disc_tengahstereo_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMBANG : {Number(
                  disc_tengahkembang_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMPIS : {Number(
                  disc_tengahkempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMBAR : {Number(
                  disc_tengahkembar_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN MONO : {Number(
                  disc_depanmono_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
                DEPAN STEREO : {Number(
                  disc_depanstereo_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN KEMBANG : {Number(
                  disc_depankembang_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN KEMPIS : {Number(
                  disc_depankempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN KEMBAR : {Number(
                  disc_depankembar_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
              </td>
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >KEI</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
              >
                BELAKANG MONO : {Number(
                  kei_belakangmono_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG STEREO : {Number(
                  kei_belakangstereo_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG KEMBANG : {Number(
                  kei_belakangkembang_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
                BELAKANG KEMPIS : {Number(
                  kei_belakangkempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                BELAKANG KEMBAR : {Number(
                  kei_belakangkembar_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH MONO : {Number(
                  kei_tengahmono_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH STEREO : {Number(
                  kei_tengahstereo_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMBANG : {Number(
                  kei_tengahkembang_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMPIS : {Number(
                  kei_tengahkempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                TENGAH KEMBAR : {Number(
                  kei_tengahkembar_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN MONO : {Number(
                  kei_depanmono_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
                DEPAN STEREO : {Number(
                  kei_depanstereo_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
                DEPAN KEMBANG : {Number(
                  kei_depankembang_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN KEMPIS : {Number(
                  kei_depankempis_bet_5050kombinasi * 100
                ).toFixed(2)}%
                <br />
                DEPAN KEMBAR : {Number(
                  kei_depankembar_bet_5050kombinasi * 100
                ).toFixed(2)}% <br />
              </td>
            </tr>
          </tbody>
        </table>
        <p
          style="padding: 5px;color:{daylight
            ? '#171717'
            : 'white'};font-size:13px;"
        >
          CARA BERMAIN
          <br />
          MONO STEREO <br />
          Menebak dari posisi Depan,Tengah,Belakang.
          <br />
          Contoh no 1234<br />
          Yang dimaksud dengan Posisi Depan adalah 2 no terdepan yaitu 12<br />
          Yang dimaksud dengan Posisi Tengah adalah 2 no ditengah yaitu 23<br />
          Yang dimaksud dengan Posisi Belakang adalah 2 no terbelakang yaitu 34<br
          />
          STEREO = Terdapat Ganjil dan Genap<br />
          MONO = Terdapat 1 pasang Ganjil atau 1 pasang Genap<br /><br />

          Analisis I : Beli Posisi Depan<br />
          Keluar : 1234<br />
          Yang menjadi pedoman adalah posisi Depan, berarti 12<br />
          12 => 1=ganjil dan 2=genap , berarti hasil = STEREO<br /><br />

          Analisis II : Beli Posisi Tengah<br />
          Keluar : 4326<br />
          Yang menjadi pedoman adalah posisi Tengah, berarti 32<br />
          32 => 3=ganjil dan 2=genap , berarti hasil = STEREO<br /><br />

          Analisis III : Beli Posisi Belakang<br />
          Keluar : 4533<br />
          Yang menjadi pedoman adalah posisi Belakang, berarti 33<br />
          33 => 3=ganjil dan 3=ganjil , berarti hasil = MONO<br />
          Jika dilakukan pembelian dengan 100rb dan menang maka: Menang = 100rb +
          [Indeks kemenangan untuk MONO STEREO]<br /><br />

          KEMBANG KEMPIS KEMBAR<br />
          Struktur=ABCD<br />
          Jika Menebang posisi Depan maka yang menjadi fokus adalah AB<br />
          Jika Menebang posisi Tengah maka yang menjadi fokus adalah BC<br />
          Jika Menebang posisi Belakang maka yang menjadi fokus adalah CD<br />
          KEMBANG jika A &#8249; B ataupun B &#8249; C ataupun C &#8249; D<br />
          KEMPIS jika A <span style="font-size:12px;">&#8250</span> B ataupun B
          <span style="font-size:12px;">&#8250</span>
          C ataupun C <span style="font-size:12px;">&#8250</span> D<br />
          KEMBAR jika A = B ataupun B = C ataupun C = D<br /><br />

          Analisis I : Beli Posisi Depan<br />
          Keluar : 4321<br />
          Yang menjadi pedoman adalah posisi Depan, berarti 43<br />
          43 => 4 <span style="font-size:12px;">&#8250</span>; 3, hasil = KEMPIS<br
          /><br />

          Analisis II : Beli Posisi Tengah<br />
          Keluar : 4236<br />
          Yang menjadi pedoman adalah posisi Tengah, berarti 23<br />
          23 => 2 &#8249; 3, hasil = KEMBANG<br /><br />

          Analisis III : Beli Posisi Belakang:<br />
          Keluar : 4099<br />
          Yang menjadi pedoman adalah posisi Belakang, berarti 99<br />
          99 => Hasil = KEMBAR<br /><br />

          Jika dilakukan pembelian dengan 100rb dan menang maka:<br />
          Menang = 100rb + [Indeks kemenangan untuk KEMBANG/KEMPIS/KEMBAR]
        </p>
      </TabPane>
    </TabContent>
  </slot:template>
</Modal>

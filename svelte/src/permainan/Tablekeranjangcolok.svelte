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
  export let count_line_colokbebas = 0;
  export let count_line_colokmacau = 0;
  export let count_line_coloknaga = 0;
  export let count_line_colokjitu = 0;
  export let min_bet_colokbebas = 0;
  export let max_bet_colokbebas = 0;
  export let disc_bet_colokbebas = 0;
  export let win_bet_colokbebas = 0;
  export let min_bet_colokmacau = 0;
  export let max_bet_colokmacau = 0;
  export let disc_bet_colokmacau = 0;
  export let win_bet_colokmacau = 0;
  export let win3_bet_colokmacau = 0;
  export let win4_bet_colokmacau = 0;
  export let min_bet_coloknaga = 0;
  export let max_bet_coloknaga = 0;
  export let disc_bet_coloknaga = 0;
  export let win_bet_coloknaga = 0;
  export let win4_bet_coloknaga = 0;
  export let min_bet_colokjitu = 0;
  export let max_bet_colokjitu = 0;
  export let disc_bet_colokjitu = 0;
  export let winas_bet_colokjitu = 0;
  export let winkop_bet_colokjitu = 0;
  export let winkepala_bet_colokjitu = 0;
  export let winekor_bet_colokjitu = 0;
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
    count_line_colokbebas +
    count_line_colokmacau +
    count_line_coloknaga +
    count_line_colokjitu;
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
        {#if group_btn_beli == true}
          <div class="col-md">
            <div class="float-end" id="btnbelitogel">
              <Button
                id="btn1"
                data-bs-toggle="modal"
                data-bs-target="#modalInformasi"
                style="margin-top:5px;"
                ><Icon name="info-circle" /> Informasi</Button
              >
              <Button
                id="btn1"
                on:click={handleRemoveKeranjang_all}
                style="margin-top:5px;"
                ><Icon name="trash" /> Hapus Semua</Button
              >
              <Button id="btn2" on:click={handleSave} style="margin-top:5px;"
                ><Icon name="cart-check" /> BELI</Button
              >
            </div>
          </div>
        {/if}
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
              {new Intl.NumberFormat().format(Math.ceil(rec.diskon))} ({Math.ceil(
                rec.diskonpercen * 100
              )}%)
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
    <h5 class="modal-title">COLOK</h5>
  </slot:template>
  <slot:template slot="body">
    <TabContent
      class="periode-menu {daylight ? '' : 'dark'}"
      style="padding: 0px;margin:0px;"
    >
      <TabPane
        tabId="informasi_colokbebas"
        tab="COLOK BEBAS"
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
                >{new Intl.NumberFormat().format(min_bet_colokbebas)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_colokbebas)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{Math.ceil(disc_bet_colokbebas * 100)}%</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win_bet_colokbebas}x</td
              >
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
          Menebak salah satu angka dari 4D. Posisi angka bisa dimana saja
          <br />
          Struktur ABCD<br /><br />

          Analisis I :<br />
          keluar : 4321<br />
          Misalnya pembelian Angka 3 dengan nilai taruhan 100rb.<br />
          Berarti menang : 100rb + [Indeks kemenangan untuk colok angka]<br
          /><br />

          Analisis I:<br />
          keluar : 4331<br />
          Misalnya pembelian Angka 3 dengan nilai taruhan 100rb.<br />
          Berarti menang: 100rb + ([Indeks kemenangan untuk colok Angka] x 2)<br
          /><br />

          Dan seterusnya untuk setiap kembaran yang berhasil ditebak, otomatis
          mendapat kelipatan [Indeks kemenangan untuk colok angka]
        </p>
      </TabPane>
      <TabPane
        tabId="informasi_colokmacau"
        tab="COLOK MACAU"
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
                >{new Intl.NumberFormat().format(min_bet_colokmacau)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_colokmacau)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{Math.ceil(disc_bet_colokmacau * 100)}%</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH 2 DIGIT</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win_bet_colokmacau}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH 3 DIGIT</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win3_bet_colokmacau}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH 4 DIGIT</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win4_bet_colokmacau}x</td
              >
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
          Cara kerja seperti colok angka tapi mesti yang keluar 2 angka dari 4D.
          <br />
          Struktur ABCD<br /><br />

          Analisis I :<br />
          keluar : 4321<br />
          Misalnya dibeli 4 dan 2 dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 2 digit]<br
          /><br />

          Analisis II :<br />
          keluar : 4321<br />
          Misalnya dibeli 4 dan 6 dengan nilai 100rb.<br />
          Berarti KALAH dan nilai betting tidak dikembalikan)<br /><br />

          Analisis III :<br />
          keluar : 4331<br />
          Misalnya dibeli 4 dan 3 dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 3 digit]<br
          /><br />

          Analisis IV :<br />
          keluar : 4334<br />
          Misalnya dibeli 4 dan 3 dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk Macau, kategori: 3 digit]
        </p>
      </TabPane>
      <TabPane
        tabId="informasi_coloknaga"
        tab="COLOK NAGA"
        style="padding:0px;font-size:12px;color:#8a8a8a;"
      >
        <table class="table table-dark">
          <tbody>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MIN BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(min_bet_coloknaga)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_coloknaga)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{Math.ceil(disc_bet_coloknaga * 100)}%</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH 3 DIGIT</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win_bet_coloknaga}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH 4 DIGIT</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{win4_bet_coloknaga}x</td
              >
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
          Cara kerja seperti colok angka tapi mesti yang keluar 3 angka dari 4D.
          <br />
          Struktur ABCD<br /><br />

          Analisis I :<br />
          keluar : 4321<br />
          Misalnya dibeli 4,2 dan 3 dengan nilai 100rb.<br />
          Berarti menang: karena keluar 3 digit,angka 4,2 dan 3.<br />
          100rb + [Indeks kemenangan untuk colok naga , kategori: 3 digit]<br
          /><br />

          Analisis II :<br />
          keluar : 4321<br />
          Misalnya dibeli 4,2 dan 6 dengan nilai 100rb.<br />
          karena keluar hanya 2 digit angka 4 dan 2, dan angka 6 tidak muncul berarti
          kalah<br /><br />

          Analisis III :<br />
          keluar : 4331<br />
          Misalnya dibeli 4,3 dan 3 dengan nilai 100rb.<br />
          Berarti menang: karena keluar 3 digit,angka 4,3 dan 3.<br />
          100rb + [Indeks kemenangan untuk colok naga , kategori: 3 digit]<br
          /><br />

          Analisis IV :<br />
          keluar : 4334<br />
          Misalnya dibeli 4,3 dan 3 dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk colok naga, kategori:
          4 digit]
        </p>
      </TabPane>
      <TabPane
        tabId="informasi_colokjitu"
        tab="COLOK JITU"
        style="padding:0px;font-size:12px;color:#8a8a8a;"
      >
        <table class="table table-dark">
          <tbody>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MIN BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(min_bet_colokjitu)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >MAX BET</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{new Intl.NumberFormat().format(max_bet_colokjitu)}</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >DISKON</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{Math.ceil(disc_bet_colokjitu * 100)}%</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH AS</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{winas_bet_colokjitu}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH KOP</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{winkop_bet_colokjitu}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH KEPALA</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{winkepala_bet_colokjitu}x</td
              >
            </tr>
            <tr>
              <td
                style="background:#303030;border:1px solid #282828;font-size:12px;"
                >HADIAH EKOR</td
              >
              <td
                style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
                >{winekor_bet_colokjitu}x</td
              >
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
          Menebak satu angka pada posisi tertentu dari 4D.
          <br />
          Struktur ABCD<br /><br />

          Analisis I :<br />
          keluar : 4321<br />
          Misalnya dibeli 4 pada posisi AS dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk colok jitu]<br /><br
          />

          Analisis II :<br />
          keluar : 4331<br />
          Misalnya dibeli 3 pada posisi KOP dengan nilai 100rb.<br />
          Berarti menang: 100rb + [Indeks kemenangan untuk colok jitu]. Hasilnya
          sama dengan analisis I karena hanya memperhatikan posisi yang dipasang.<br
          /><br />

          Analisis III :<br />
          keluar : 4331<br />
          Misalnya dibeli 4 pada posisi EKOR dengan nilai 100rb.<br />
          Berarti kalah. Biarpun nilai 4 keluar pada posisi AS tapi tidak akan mepengaruhi
          pemilihan di pososi EKOR
        </p>
      </TabPane>
    </TabContent>
  </slot:template>
</Modal>

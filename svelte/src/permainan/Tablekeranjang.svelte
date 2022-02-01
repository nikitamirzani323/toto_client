<script>
  import { Button, Icon } from "sveltestrap";
  import { createEventDispatcher } from "svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Modal from "../components/Modalfull2.svelte";

  export let keranjang = [];
  export let totalkeranjang = 0;
  export let client_device = "";
  export let group_btn_beli = false;
  export let count_line_4d = 0;
  export let count_line_3d = 0;
  export let count_line_3dd = 0;
  export let count_line_2d = 0;
  export let count_line_2dd = 0;
  export let count_line_2dt = 0;
  export let minimal_bet = 0;
  export let max4d_bet = 0;
  export let max3d_bet = 0;
  export let max3dd_bet = 0;
  export let max2d_bet = 0;
  export let max2dd_bet = 0;
  export let max2dt_bet = 0;
  export let disc4d_bet = 0;
  export let disc3d_bet = 0;
  export let disc3dd_bet = 0;
  export let disc2d_bet = 0;
  export let disc2dd_bet = 0;
  export let disc2dt_bet = 0;
  export let win4d_bet = 0;
  export let win3d_bet = 0;
  export let win3dd_bet = 0;
  export let win2d_bet = 0;
  export let win2dd_bet = 0;
  export let win2dt_bet = 0;
  export let win4dnodiskon_bet = 0;
  export let win3dnodiskon_bet = 0;
  export let win3ddnodiskon_bet = 0;
  export let win2dnodiskon_bet = 0;
  export let win2ddnodiskon_bet = 0;
  export let win2dtnodiskon_bet = 0;
  export let win4dbb_kena_bet = 0;
  export let win3dbb_kena_bet = 0;
  export let win3ddbb_kena_bet = 0;
  export let win2dbb_kena_bet = 0;
  export let win2ddbb_kena_bet = 0;
  export let win2dtbb_kena_bet = 0;
  export let win4dbb_bet = 0;
  export let win3dbb_bet = 0;
  export let win3ddbb_bet = 0;
  export let win2dbb_bet = 0;
  export let win2ddbb_bet = 0;
  export let win2dtbb_bet = 0;
  export let limitline_4d = 0;
  export let limitline_3d = 0;
  export let limitline_3dd = 0;
  export let limitline_2d = 0;
  export let limitline_2dd = 0;
  export let limitline_2dt = 0;
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
    count_line_4d +
    count_line_3d +
    count_line_3dd +
    count_line_2d +
    count_line_2dd +
    count_line_2dt;
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
            NOWRAP>TIPE</th
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
              style="text-align:center;vertical-align:top;"
              on:click={() => {
                handleRemoveKeranjang(rec.id, rec.bayar);
              }}
            >
              <Icon name="trash" style="cursor:pointer;" />
            </td>
            <td class="table-fill" class:dark={daylight === false}
              >{rec.nomor}</td
            >
            <td class="table-fill" class:dark={daylight === false}
              >{rec.tipetoto}</td
            >
            <td class="table-fill" class:dark={daylight === false}
              >{rec.permainan}</td
            >
            <td class="table-fill" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(rec.bet)}
            </td>
            <td class="table-fill" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(Math.ceil(rec.diskon))} ({(
                rec.diskonpercen * 100
              ).toFixed(1)}%)
            </td>
            <td class="table-fill" class:dark={daylight === false}>
              {new Intl.NumberFormat().format(rec.bayar)}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <div class="container mt-3">
      <div class="row">
        <div class="col-md fs-4">
          TOTAL LINE : <span class="footer-font" class:dark={daylight === false}
            >{count_line}</span
          >
        </div>
      </div>
      <div class="row">
        <div class="col fs-6">
          4D : <span class="footer-font" class:dark={daylight === false}
            >{count_line_4d}</span
          >
        </div>
        <div class="col fs-6">
          3D : <span class="footer-font" class:dark={daylight === false}
            >{count_line_3d}</span
          >
        </div>
        <div class="col fs-6">
          3DD : <span class="footer-font" class:dark={daylight === false}
            >{count_line_3dd}</span
          >
        </div>
        <div class="col fs-6">
          2D : <span class="footer-font" class:dark={daylight === false}
            >{count_line_2d}</span
          >
        </div>
        <div class="col fs-6">
          2DD : <span class="footer-font" class:dark={daylight === false}
            >{count_line_2dd}</span
          >
        </div>
        <div class="col fs-6">
          2DT : <span class="footer-font" class:dark={daylight === false}
            >{count_line_2dt}</span
          >
        </div>
      </div>
    </div>
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
    <h5 class="modal-title">4D, 3D dan 2D</h5>
  </slot:template>
  <slot:template slot="body">
    <table class="table table-dark">
      <thead>
        <tr>
          <th
            style="background:#303030;border:1px solid #282828;border-bottom:none;"
            >&nbsp;</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >4D</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >3D</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >3DD</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >2D</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >2DD</th
          >
          <th
            style="background:#303030;border:1px solid #282828;text-align:right;border-bottom:none;"
            >2DT</th
          >
        </tr>
      </thead>
      <tbody>
        <tr>
          <td
            nowrap
            style="background:#303030;border:1px solid #282828;font-size:12px;"
            >MIN BET</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(minimal_bet)}</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(minimal_bet)}</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(minimal_bet)}</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(minimal_bet)}</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(minimal_bet)}</td
          >
          <td
            nowrap
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{new Intl.NumberFormat().format(minimal_bet)}</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >MAX BET</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max4d_bet)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max3d_bet)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max3dd_bet)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max2d_bet)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max2dd_bet)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max2dt_bet)}</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >DISKON</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc4d_bet * 100)}%</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc3d_bet * 100)}%</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc3dd_bet * 100)}%</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc2d_bet * 100)}%</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc2dd_bet * 100)}%</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{Math.ceil(disc2dt_bet * 100)}%</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >HADIAH</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{win4d_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3d_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3dd_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2d_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dd_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dt_bet}x</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >HADIAH FULL</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{win4dnodiskon_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3dnodiskon_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3ddnodiskon_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dnodiskon_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2ddnodiskon_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dtnodiskon_bet}x</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >HADIAH BB KENA</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{win4dbb_kena_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3dbb_kena_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3ddbb_kena_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dbb_kena_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2ddbb_kena_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dtbb_kena_bet}x</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >HADIAH BB</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{win4dbb_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3dbb_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win3ddbb_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dbb_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2ddbb_bet}x</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{win2dtbb_bet}x</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:11px;"
            >LIMIT LINE</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(limitline_4d)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(limitline_3d)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(limitline_3dd)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
          >
            {new Intl.NumberFormat().format(limitline_2d)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(limitline_2dd)}</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:11px;color:#ffd292;"
            >{new Intl.NumberFormat().format(limitline_2dt)}</td
          >
        </tr>
      </tbody>
    </table>
    <p style="font-size:13px;padding:5px;color:white;">
      <b>Contoh Penulisan 4D/3D/2D:</b><br />
      1234 : 4D<br />
      123 : 3D<br /> *123 : 3D<br /> 123* : 3DD<br />
      12 : 2D <br />
      **12 : 2D<br /> 12** : 2DD<br /> *12* : 2DT
    </p>
    <p style="font-size:13px;padding: 5px;color:white;">
      <b>CARA BERMAIN</b><br />
      Menebak 4 angka, 3 angka dan 2 angka
      <br />
      Struktur ABCD<br /><br />

      Misalnya keluar : 4321<br />
      Berarti pemenang untuk<br />
      4D = 4321<br />
      3D = 321<br />
      3DD = 432<br />
      2D = 21<br />
      2DD = 43<br />
      2DT = 32<br /><br />

      Aturan permainan:<br />
      1. Jika anda membeli diluar dari nomor yang dikeluarkan, berarti anda kalah
      dan uang tidak dikembalikan sama sekali.<br />
      2. Jika anda membeli masing 100rb untuk angka :<br />
      4D = 4321<br />
      3D = 321<br />
      2D = 21<br /><br />
      Berarti kemenangan anda adalah :<br />
      4D = 100rb x [Indeks kemenangan untuk 4D]<br />
      3D = 100rb x [Indeks kemenangan untuk 3D]<br />
      2D = 100rb x [Indeks kemenangan untuk 2D]<br /><br />
      (Catatan: nilai bet 100rb tidak dikembalikan ke member)<br />
      (Khusus untuk 4D,3D dan 2D diberikan diskon tambahan)<br />
    </p>
  </slot:template>
</Modal>

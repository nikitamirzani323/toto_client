<script>
  import { Button, Icon } from "sveltestrap";
  import { createEventDispatcher } from "svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Modal from "../components/Modalfull2.svelte";
  import Lineinfo from "../components/Keranjang.svelte";

  export let keranjang = [];
  export let client_device = "";
  export let totalkeranjang = 0;
  export let group_btn_beli = false;
  export let count_line_dasar = 0;
  export let count_line_standart = 0;
  export let min_bet = 0;
  export let max_bet = 0;
  export let win_bet = 0;
  export let diskon_bet = 0;
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
  $: count_line = count_line_dasar + count_line_standart;
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
    <h5 class="modal-title">SHIO</h5>
  </slot:template>
  <slot:template slot="body">
    <table class="table table-dark table-sm">
      <tbody>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:12px;"
            >MIN BET</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{new Intl.NumberFormat().format(min_bet)}</td
          >
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:12px;"
            >MAX BET</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
            >{new Intl.NumberFormat().format(max_bet)}</td
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
            {(diskon_bet * 100).toFixed(1)}%
          </td>
        </tr>
        <tr>
          <td
            style="background:#303030;border:1px solid #282828;font-size:12px;"
            >HADIAH</td
          >
          <td
            style="border:1px solid #282828;text-align:right;font-size:12px;color:#ffd292;"
          >
            {win_bet.toFixed(1)}X
          </td>
        </tr>
      </tbody>
    </table>
    <p style="font-size:13px;padding: 5px;color:white;">
      <b>CARA BERMAIN</b>
      <br />
      Menebak SHIO dari posisi 2D, SHIO merupakan 12 lambang kelahiran dalam penanggalan
      China. Dalam permainan ini, setiap lambang diwakili dengan satu nomor.
      <br />
      Struktur ABCD<br /><br />

      Analisis I :<br />
      Keluar : 4321<br />
      Permainan ini hanya memperhatikan posisi 2D, berarti yang dipedomanin = 21<br
      />
      Hasil = 21-12 = 9 (shio disesuaikan dengan tabel diatas)<br />
      catatan: nilai yang dikurangi merupakan kelipatan 12.<br /><br />

      Jika dilakukan pembelian dengan 100rb dan menang maka:<br />
      Menang = 100rb + [Indeks kemenangan untuk SHIO]<br /><br />

      NB: Indeks menang dan diskon dapat dilihat di bagian Peraturan
    </p>
  </slot:template>
</Modal>

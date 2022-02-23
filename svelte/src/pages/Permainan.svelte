<script>
  import { Button, Col } from "sveltestrap";
  import dayjs from "dayjs";
  import Loader2 from "../components/Loader.svelte";
  import Notif from "../components/Notif.svelte";
  import Headerback from "../components/Header.svelte";
  import Panel from "../components/Panel.svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Placeholder from "../components/Placeholder.svelte";
  import Form432d from "../permainan/Form432d.svelte";
  import Formcolok from "../permainan/Formcolok.svelte";
  import Form5050 from "../permainan/Form5050.svelte";
  import Formkombinasi from "../permainan/Formkombinasi.svelte";
  import Formdasar from "../permainan/Formdasar.svelte";
  import Formshio from "../permainan/Formshio.svelte";
  import Fa from "svelte-fa";
  import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";
  import Swal from "sweetalert2";

  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_credit = 0;
  export let client_ipaddress = "";
  export let client_timezone = "Asia/Jakarta";
  export let client_device = "";
  export let pasaran_code = "";
  export let pasaran_name = "";
  export let pasaran_periode = 0;
  export let permainan = "";
  export let revisi_note = "";
  export let revisi_periode = "";
  export let revisi_pasaran = "";

  export let daylight = false;
  export let checked;
  export let balance_credit;
  export let agent_home_url = "";
  let home = false;
  let css_loader = "display:none;";

  let resultinvoice = [];
  let resultslipbet = [];
  let filterBukuMimpi = [];
  let listBukumimpi = [];
  let record = "";
  let idcomppasaran = "";
  let idtrxkeluaran = "";
  let statuspasaran = "";
  let permainan_title = "4D / 3D / 2D";
  let permainan_periode = pasaran_periode + " - " + pasaran_code;
  let totalbayar_invoice = 0;
  let totalbet_invoice = 0;
  let searchbukumimpi = "";
  let tipe = "";
  let message_err = "";
  let css_err = "display:none;";
  let tabs = ["4-3-2", "colok", "5050", "kombinasi", "dasar", "shio"];
  let selected = tabs[0];

  const changePermainan = (e) => {
    permainan = e;
    switch (e) {
      case "4-3-2":
        permainan_title = "4D / 3D / 2D";
        selected = tabs[0];
        break;
      case "colok":
        permainan_title = "COLOK";
        selected = tabs[1];
        break;
      case "5050":
        permainan_title = "50 - 50";
        selected = tabs[2];
        break;
      case "kombinasi":
        permainan_title = "MACAU / KOMBINASI";
        selected = tabs[3];
        break;
      case "dasar":
        permainan_title = "DASAR";
        selected = tabs[4];
        break;
      case "shio":
        permainan_title = "SHIO";
        selected = tabs[5];
        break;
    }
  };

  async function checkpasaran() {
    css_loader =
      "position:absolute;z-index: 1000;left: 0%;top: 35%;display:inline;";
    const res = await fetch("/api/checkpasaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        timezone: client_timezone,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json;
    if (json.status == "200") {
      css_loader = "display:none;";
    }
    idcomppasaran = record["pasaran_idcomp"];
    idtrxkeluaran = record["pasaran_idtransaction"];
    statuspasaran = record["pasaran_status"];
    pasaran_periode = record["pasaran_periode"];
    pasaran_name = record["pasaran_name"];
    if (statuspasaran == "OFFLINE") {
      css_err = "display:inline-block";
      message_err = `Pasaran ${pasaran_name} sedang OFFLINE`;
      Swal.fire({
        title: message_err,
        icon: "error",
        confirmButtonText: "Go back",
        confirmButtonColor: "#00a86b",
        heightAuto: false,
        allowOutsideClick: false,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
      }).then((result) => {
        if (result.isConfirmed) {
          location.href = agent_home_url;
        }
      });
      // setTimeout(function () {
      //   css_err = "display: none;";
      // }, 5000);
    }
    invoicebet("all");
  }

  async function checkrevisi() {
    const res = await fetch("/api/checkrevisi", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        timezone: client_timezone,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json;
    if (json.status == "200") {
      css_loader = "display:none;";
      revisi_note = record["pasaran_revisi"];
      revisi_periode = record["pasaran_periode"];
      revisi_pasaran = record["pasaran_name"];
    }
  }

  async function invoicebet(e) {
    const res = await fetch("/api/invoicebet", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        username: client_username,
        pasaran_code: pasaran_code,
        pasaran_periode: pasaran_periode,
        permainan: e,
        pasaran_idtransaction: parseInt(idtrxkeluaran),
      }),
    });

    const json = await res.json();
    record = json.record;

    if (json.status == 200) {
      totalbayar_invoice = json.totalbayar;
      totalbet_invoice = json.totalbet;
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          resultinvoice = [
            ...resultinvoice,
            {
              permainan: record[i]["permainan"],
              periode: record[i]["periode"],
              nomor: record[i]["nomor"],
              bet: record[i]["bet"],
              diskon: record[i]["diskon"],
              kei: record[i]["kei"],
              bayar: record[i]["bayar"],
              win: record[i]["win"],
              menang: record[i]["menang"],
            },
          ];
        }
      }
    }
  }
  async function slipbet() {
    const res = await fetch("/api/slipperiode", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        username: client_username,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json.record;

    if (json.status == 200) {
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          resultslipbet = [
            ...resultslipbet,
            {
              status: record[i]["status"],
              tglkeluaran: record[i]["tglkeluaran"],
              idinvoice: record[i]["idinvoice"],
              periode: record[i]["periode"],
              totallose: record[i]["totallose"],
              backgroundstatus: record[i]["background"],
            },
          ];
        }
      }
    }
  }
  async function fetch_bukumimpi() {
    const res = await fetch("/api/bookdream", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        tipe: tipe,
        nama: searchbukumimpi.toLowerCase(),
      }),
    });

    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          listBukumimpi = [
            ...listBukumimpi,
            {
              bukumimpi_tipe: record[i]["bukumimpi_type"],
              bukumimpi_nama: record[i]["bukumimpi_name"],
              bukumimpi_nomor: record[i]["bukumimpi_nomor"],
            },
          ];
        }
      } else {
        Swal.fire({
          position: "center",
          icon: "error",
          title: "Pencarian tidak ditemukan",
          showConfirmButton: false,
          timer: 3000,
          background: daylight ? "#fff" : "#171717",
          color: daylight ? "#00a86b" : "#ff9900",
          toast: true,
        });
      }
    } else {
      Swal.fire({
        position: "center",
        icon: "error",
        title: "Terjadi kesalahan pada server",
        showConfirmButton: false,
        timer: 3000,
        background: daylight ? "#fff" : "#171717",
        color: daylight ? "#00a86b" : "#ff9900",
        toast: true,
      });
    }
  }
  const handleInvoice = (e) => {
    resultinvoice = [];
    resultslipbet = [];
    invoicebet("all");
    slipbet();
  };
  checkpasaran();
  checkrevisi();

  const handleSelect = (event) => {
    changePermainan(event.target.value);
  };
  const handleClickTab = (e) => {
    switch (e) {
      case "SLIP":
        resultslipbet = [];
        slipbet();
        break;
      case "BUKUMIMPI":
        filterBukuMimpi = [];
        listBukumimpi = [];
        searchbukumimpi = "";
        tipe = "";
        fetch_bukumimpi();
        break;
    }
  };
  const handleClickBukuMimpi = (e) => {
    filterBukuMimpi = [];
    listBukumimpi = [];
    switch (e) {
      case "ALL":
        tipe = "";
        searchbukumimpi = "";
        break;
      case "4D":
        tipe = "4D";
        searchbukumimpi = "";
        break;
      case "3D":
        tipe = "3D";
        searchbukumimpi = "";
        break;
      case "2D":
        tipe = "2D";
        searchbukumimpi = "";
        break;
    }
    fetch_bukumimpi();
  };
  const handleKeyboardbukumimpi_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      // tipe = "";
      filterBukuMimpi = [];
      listBukumimpi = [];
      fetch_bukumimpi();
    }
  };

  $: {
    if (searchbukumimpi) {
      filterBukuMimpi = listBukumimpi.filter((item) =>
        item.bukumimpi_nama
          .toLowerCase()
          .includes(searchbukumimpi.toLowerCase())
      );
    } else {
      filterBukuMimpi = [...listBukumimpi];
    }
  }
</script>

{#if statuspasaran == "ONLINE"}
  <Headerback
    {client_username}
    {client_company}
    {client_credit}
    {client_token}
    {client_ipaddress}
    {client_timezone}
    {client_device}
    bind:checked
    {revisi_note}
    {revisi_periode}
    {revisi_pasaran}
    {daylight}
    {home}
  />
  {#if client_device == "WEBSITE"}
    <Col xs="12" style="padding:0;">
      <Panel {daylight}>
        <slot:template slot="body">
          <nav class="navbar">
            {#if agent_home_url !== ""}
              <a class="navbar-brand branding" href={agent_home_url}
                ><Fa icon={faArrowLeft} size="1x" /> Back</a
              >
            {:else}
              <a
                class="navbar-brand branding"
                href="/?token={client_token}&agent={client_company}"
                ><Fa icon={faArrowLeft} size="1x" /> Back</a
              >
            {/if}
            <ul
              class="nav top-menu nav-pills justify-content-center"
              class:dark={daylight === false}
            >
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "4-3-2"}
                  on:click={() => {
                    changePermainan("4-3-2");
                  }}>4D/3D/2D</a
                >
              </li>
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "colok"}
                  on:click={() => {
                    changePermainan("colok");
                  }}>Colok</a
                >
              </li>
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "5050"}
                  on:click={() => {
                    changePermainan("5050");
                  }}>50-50</a
                >
              </li>
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "kombinasi"}
                  on:click={() => {
                    changePermainan("kombinasi");
                  }}>KOMBINASI</a
                >
              </li>
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "dasar"}
                  on:click={() => {
                    changePermainan("dasar");
                  }}>DASAR</a
                >
              </li>
              <li class="nav-item">
                <!-- svelte-ignore a11y-missing-attribute -->
                <a
                  class="nav-link"
                  class:active={selected === "shio"}
                  on:click={() => {
                    changePermainan("shio");
                  }}>SHIO</a
                >
              </li>
            </ul>
            <div class="justify-content-center" />
          </nav>
        </slot:template>
      </Panel>
    </Col>
    <div class="clearfix mt-2" />
    <Col
      xxl="7"
      xl="7"
      lg="7"
      md="7"
      sm="12"
      xs="12"
      style="padding:0px;padding-right:2px;"
    >
      {#if permainan == "4-3-2"}
        <Form432d
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {daylight}
          bind:balance_credit
        />
      {/if}
      {#if permainan == "colok"}
        <Formcolok
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
          {daylight}
          bind:balance_credit
        />
      {/if}
      {#if permainan == "5050"}
        <Form5050
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
          {daylight}
          bind:balance_credit
        />
      {/if}
      {#if permainan == "kombinasi"}
        <Formkombinasi
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
          {daylight}
          bind:balance_credit
        />
      {/if}
      {#if permainan == "dasar"}
        <Formdasar
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
          {daylight}
          bind:balance_credit
        />
      {/if}
      {#if permainan == "shio"}
        <Formshio
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
          {daylight}
          bind:balance_credit
        />
      {/if}
    </Col>
    <Col
      xxl="5"
      xl="5"
      lg="5"
      md="5"
      sm="12"
      xs="12"
      style="padding:0px;padding-left:2px;"
    >
      <ul
        class="nav nav-pills bet-pannel mb-3"
        class:dark={daylight === false}
        id="pills-tab"
        role="tablist"
      >
        <li class="nav-item" role="presentation">
          <button
            class="nav-link active"
            id="pills-home-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-pasangan"
            type="button"
            role="tab"
            aria-controls="pills-pasangan"
            aria-selected="true">BET HISTORY</button
          >
        </li>

        <li
          on:click={() => {
            handleClickTab("BUKUMIMPI");
          }}
          class="nav-item"
          role="presentation"
        >
          <button
            class="nav-link"
            id="pills-contact-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-bukumimpi"
            type="button"
            role="tab"
            aria-controls="pills-bukumimpi"
            aria-selected="false">BUKU MIMPI</button
          >
        </li>
      </ul>
      <div class="tab-content" id="pills-tabContent">
        <div
          class="tab-pane fade show active"
          id="pills-pasangan"
          role="tabpanel"
          aria-labelledby="pills-pasangan-tab"
        >
          <PanelFull
            header={true}
            footer={true}
            header_style="padding:0px;margin:0px;"
            body_style="padding:0px;margin:0px;background: url(/bg-invoice.svg);
            background-position: bottom right;
            background-repeat: no-repeat;
            background-size: 400px;
            border:1px solid {daylight ? '#fff' : '#0e0c13'};height:655px;"
            {daylight}
          >
            <slot:template slot="header">
              <div
                class="row"
                style="padding: 20px;font-size: 15px;color: {daylight
                  ? '#171717'
                  : '#fff'};"
              >
                <div class="col">
                  <span
                    >BET SAYA : <span
                      style="color:{daylight
                        ? '#00a86b'
                        : '#FF9900'}; font-weight:bold"
                      >{new Intl.NumberFormat().format(totalbet_invoice)}</span
                    ></span
                  >
                </div>
                <div class="col text-end">
                  <span
                    >TOTAL BAYAR : <span
                      style="color:{daylight
                        ? '#00a86b'
                        : '#FF9900'}; font-weight:bold"
                      >{new Intl.NumberFormat().format(
                        totalbayar_invoice
                      )}</span
                    ></span
                  >
                </div>
              </div>
            </slot:template>
            <slot:template slot="body">
              <table
                class="table"
                class:table-dark={daylight === false}
                class:table-striped={daylight === false}
              >
                <thead>
                  <tr>
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;font-size:13px;border-bottom:none;"
                      NOWRAP>NOMOR</th
                    >
                    <th
                      width="10%"
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
                      NOWRAP>KEI(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
                      NOWRAP>DIS(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;font-size:13px;border-bottom:none;"
                      NOWRAP>BAYAR</th
                    >
                  </tr>
                </thead>
                <tbody>
                  {#each resultinvoice as rec}
                    <tr>
                      <td
                        class="table-fill text-center"
                        class:dark={daylight === false}
                        NOWRAP>{rec.nomor}</td
                      >
                      <td
                        class="table-fill text-center"
                        class:dark={daylight === false}
                        NOWRAP>{rec.permainan}</td
                      >
                      <td
                        class="table-fill text-end"
                        class:dark={daylight === false}
                        NOWRAP>{new Intl.NumberFormat().format(rec.bet)}</td
                      >
                      <td
                        class="table-fill text-end"
                        class:dark={daylight === false}
                        NOWRAP>{rec.kei.toFixed(1)}</td
                      >
                      <td
                        class="table-fill text-end"
                        class:dark={daylight === false}
                        NOWRAP>{rec.diskon.toFixed(1)}</td
                      >
                      <td
                        class="table-fill text-end"
                        class:dark={daylight === false}
                        NOWRAP>{new Intl.NumberFormat().format(rec.bayar)}</td
                      >
                    </tr>
                  {/each}
                </tbody>
              </table>
            </slot:template>
            <slot:template slot="footer" />
          </PanelFull>
        </div>
        <div
          class="tab-pane fade"
          id="pills-bukumimpi"
          role="tabpanel"
          aria-labelledby="pills-result-tab"
        >
          <div
            class="tab-content periode-menu"
            class:dark={daylight === false}
            id="pills-tabContent"
          >
            <ul
              class="nav nav-tabs"
              id="pills-tab"
              role="tablist"
              style="background-color: {daylight
                ? '#fff'
                : '#181818'};overflow-x:hidden;"
            >
              <li
                on:click={() => {
                  handleClickBukuMimpi("ALL");
                }}
                class="nav-item"
                role="presentation"
              >
                <button
                  class="nav-link active"
                  id="pills-home-tab"
                  data-bs-toggle="pill"
                  data-bs-target="#pills-bukumimpiall"
                  type="button"
                  role="tab"
                  aria-controls="pills-bukumimpiall"
                  aria-selected="true">ALL</button
                >
              </li>
              <li
                on:click={() => {
                  handleClickBukuMimpi("4D");
                }}
                class="nav-item"
                role="presentation"
              >
                <button
                  class="nav-link"
                  id="pills-profile-tab"
                  data-bs-toggle="pill"
                  data-bs-target="#pills-bukumimpi4d"
                  type="button"
                  role="tab"
                  aria-controls="pills-bukumimpi4d"
                  aria-selected="false">4D</button
                >
              </li>
              <li
                on:click={() => {
                  handleClickBukuMimpi("3D");
                }}
                class="nav-item"
                role="presentation"
              >
                <button
                  class="nav-link"
                  id="pills-contact-tab"
                  data-bs-toggle="pill"
                  data-bs-target="#pills-bukumimpi3d"
                  type="button"
                  role="tab"
                  aria-controls="pills-bukumimpi3d"
                  aria-selected="false">3D</button
                >
              </li>
              <li
                on:click={() => {
                  handleClickBukuMimpi("2D");
                }}
                class="nav-item"
                role="presentation"
              >
                <button
                  class="nav-link"
                  id="pills-contact-tab"
                  data-bs-toggle="pill"
                  data-bs-target="#pills-bukumimpi2d"
                  type="button"
                  role="tab"
                  aria-controls="pills-bukumimpi2d"
                  aria-selected="false">2D</button
                >
              </li>
            </ul>
            <div class="form-floating">
              <input
                bind:value={searchbukumimpi}
                on:keypress={handleKeyboardbukumimpi_checkenter}
                style="border-radius: none;border: none; background: {daylight
                  ? '#fff'
                  : '#303030'} none repeat scroll 0% 0%; color: {daylight
                  ? '#171717'
                  : '#fff'}; font-size: 15px; "
                placeholder="Ketik apa yang telah kamu impikan"
                class="form-control button-bet-default"
                type="text"
                id="cariMimpi"
                autocomplete="off"
              />
              <label for="cariMimpi" class="form-label"
                >Ketik apa yang telah kamu impikan</label
              >
            </div>

            <div
              class="tab-pane fade show active"
              id="pills-bukumimpiall"
              role="tabpanel"
              aria-labelledby="pills-bukumimpiall-tab"
            >
              <PanelFull
                header={false}
                footer={false}
                body_style="padding:0px;margin:0px;height:650px;"
                {daylight}
              >
                <slot:template slot="body">
                  {#if filterBukuMimpi !== ""}
                    <table>
                      <tbody>
                        {#each filterBukuMimpi as rec}
                          <tr>
                            <td
                              NOWRAP
                              width="30px"
                              style="text-align:center;vertical-align:top;font-size:14px;"
                              class="color-font"
                              class:dark={daylight === false}
                              >{rec.bukumimpi_tipe}</td
                            >
                            <td
                              width="*"
                              style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                              >{rec.bukumimpi_nama}
                              <br />
                              <span
                                style="font-size:14px;"
                                class="color-font"
                                class:dark={daylight === false}
                                >{rec.bukumimpi_nomor}</span
                              >
                            </td>
                          </tr>
                        {/each}
                      </tbody>
                    </table>
                  {:else}
                    <Placeholder
                      total_placeholder="20"
                      card_style="background-color:#2c2c2c;border:none;margin-top:5px;"
                    />
                  {/if}
                </slot:template>
              </PanelFull>
            </div>
            <div
              class="tab-pane fade"
              id="pills-bukumimpi4d"
              role="tabpanel"
              aria-labelledby="pills-bukumimpi4d-tab"
            >
              <PanelFull
                header={false}
                footer={false}
                body_style="padding:0px;margin:0px;height:650px;"
                {daylight}
              >
                <slot:template slot="body">
                  {#if filterBukuMimpi !== ""}
                    <table>
                      <tbody>
                        {#each filterBukuMimpi as rec}
                          <tr>
                            <td
                              NOWRAP
                              width="30px"
                              style="text-align:center;vertical-align:top;font-size:14px;"
                              class="color-font"
                              class:dark={daylight === false}
                              >{rec.bukumimpi_tipe}</td
                            >
                            <td
                              width="*"
                              style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                              >{rec.bukumimpi_nama}
                              <br />
                              <span
                                style="font-size:14px;"
                                class="color-font"
                                class:dark={daylight === false}
                                >{rec.bukumimpi_nomor}</span
                              >
                            </td>
                          </tr>
                        {/each}
                      </tbody>
                    </table>
                  {:else}
                    <Placeholder
                      total_placeholder="20"
                      card_style="background-color:#2c2c2c;border:none;margin-top:5px;"
                    />
                  {/if}
                </slot:template>
              </PanelFull>
            </div>
            <div
              class="tab-pane fade"
              id="pills-bukumimpi3d"
              role="tabpanel"
              aria-labelledby="pills-bukumimpi3d-tab"
            >
              <PanelFull
                header={false}
                footer={false}
                body_style="padding:0px;margin:0px;height:650px;"
                {daylight}
              >
                <slot:template slot="body">
                  {#if filterBukuMimpi != ""}
                    <table>
                      <tbody>
                        {#each filterBukuMimpi as rec}
                          <tr>
                            <td
                              NOWRAP
                              width="30px"
                              style="text-align:center;vertical-align:top;font-size:14px;"
                              class="color-font"
                              class:dark={daylight === false}
                              >{rec.bukumimpi_tipe}</td
                            >
                            <td
                              width="*"
                              style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                              class="color-font"
                              class:dark={daylight === false}
                              >{rec.bukumimpi_nama}
                              <br />
                              <span style="font-size:14px;"
                                >{rec.bukumimpi_nomor}</span
                              >
                            </td>
                          </tr>
                        {/each}
                      </tbody>
                    </table>
                  {:else}
                    <Placeholder
                      total_placeholder="20"
                      card_style="background-color:#2c2c2c;border:none;margin-top:5px;"
                    />
                  {/if}
                </slot:template>
              </PanelFull>
            </div>
            <div
              class="tab-pane fade"
              id="pills-bukumimpi2d"
              role="tabpanel"
              aria-labelledby="pills-bukumimpi2d-tab"
            >
              <PanelFull
                header={false}
                footer={false}
                body_style="padding:0px;margin:0px;height:650px;"
                {daylight}
              >
                <slot:template slot="body">
                  {#if filterBukuMimpi !== ""}
                    <table>
                      <tbody>
                        {#each filterBukuMimpi as rec}
                          <tr>
                            <td
                              NOWRAP
                              width="30px"
                              style="text-align:center;vertical-align:top;font-size:14px;"
                              class="color-font"
                              class:dark={daylight === false}
                              >{rec.bukumimpi_tipe}</td
                            >
                            <td
                              width="*"
                              style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                              >{rec.bukumimpi_nama}
                              <br />
                              <span
                                style="font-size:14px;"
                                class="color-font"
                                class:dark={daylight === false}
                                >{rec.bukumimpi_nomor}</span
                              >
                            </td>
                          </tr>
                        {/each}
                      </tbody>
                    </table>
                  {:else}
                    <Placeholder
                      total_placeholder="20"
                      card_style="background-color:#2c2c2c;border:none;margin-top:5px;"
                    />
                  {/if}
                </slot:template>
              </PanelFull>
            </div>
          </div>
        </div>
      </div>
    </Col>
  {:else}
    <div style="margin-top:5px;">&nbsp;</div>

    <div class="tab-content" id="pills-tabContent">
      <ul
        class="nav nav-fill nav-pills mb-3"
        id="pills-tab"
        role="tablist"
        style="border:1px solid #f3f3f3; border-radius:4px; background-color: {daylight
          ? '#fff'
          : '#181818'};"
      >
        <li class="nav-item" role="presentation">
          <button
            style="font-size:10px;"
            class="nav-link custom active"
            class:custom-dark={daylight === false}
            id="pills-keranjang-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-keranjang"
            type="button"
            role="tab"
            aria-controls="pills-keranjang"
            aria-selected="true">KERANJANG</button
          >
        </li>
        <li class="nav-item" role="presentation">
          <button
            style="font-size:10px;"
            class="nav-link custom"
            class:custom-dark={daylight === false}
            id="pills-pasangan-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-pasangan"
            type="button"
            role="tab"
            aria-controls="pills-pasangan"
            aria-selected="true">BET HISTORY</button
          >
        </li>
      </ul>
      <div
        class="tab-pane fade show active"
        id="pills-keranjang"
        role="tabpanel"
        aria-labelledby="pills-keranjang-tab"
      >
        <br />
        <b style="font-size: 13px;">Pilih Permainan Dibawah Ini : </b>
        <!-- svelte-ignore a11y-no-onchange -->
        <select
          on:change={handleSelect}
          style="background-color: {daylight
            ? '#fff'
            : '#323030'};color: {daylight
            ? '#171717'
            : '#fff'};border:1px solid {daylight ? '#f3f3f3' : '#323030'};"
          aria-label="Permainan"
          class="form-select"
        >
          <option value="4-3-2">4D/3D/2D</option>
          <option value="colok">COLOK</option>
          <option value="5050">50-50</option>
          <option value="kombinasi">MACAU / KOMBINASI</option>
          <option value="dasar">DASAR</option>
          <option value="shio">SHIO</option>
        </select>
        <div class="clear-fix" />
        <br />
        {#if permainan == "4-3-2"}
          <Form432d
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
        {#if permainan == "colok"}
          <Formcolok
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
        {#if permainan == "5050"}
          <Form5050
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
        {#if permainan == "kombinasi"}
          <Formkombinasi
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
        {#if permainan == "dasar"}
          <Formdasar
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
        {#if permainan == "shio"}
          <Formshio
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
            {daylight}
            bind:balance_credit
          />
        {/if}
      </div>
      <div
        class="tab-pane"
        id="pills-pasangan"
        role="tabpanel"
        aria-labelledby="pills-pasangan-tab"
      >
        <PanelFull
          header={true}
          footer={true}
          body_style="padding:0px;margin:0px;background: url(/bg-invoice.svg);
          background-position: bottom right;
          background-repeat: no-repeat;
          background-size: 400px;
          border:1px solid {daylight ? '#fff' : '#0e0c13'};height:450px;"
          {daylight}
        >
          <slot:template slot="header">
            <div
              class="row"
              style="padding: 5px;font-size: 15px;color: {daylight
                ? '#171717'
                : '#fff'};"
            >
              <div class="col">
                <span
                  >BET SAYA : <span
                    style="color:{daylight
                      ? '#00a86b'
                      : '#FF9900'}; font-weight:bold"
                    >{new Intl.NumberFormat().format(totalbet_invoice)}</span
                  ></span
                >
              </div>
              <div class="col text-end">
                <span
                  >TOTAL BAYAR : <span
                    style="color:{daylight
                      ? '#00a86b'
                      : '#FF9900'}; font-weight:bold"
                    >{new Intl.NumberFormat().format(totalbayar_invoice)}</span
                  ></span
                >
              </div>
            </div>
          </slot:template>
          <slot:template slot="body">
            <table
              class="table"
              class:table-dark={daylight === false}
              class:table-striped={daylight === false}
            >
              <thead>
                <tr>
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>NOMOR</th
                  >
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>PERMAINAN</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>BET</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>KEI(%)</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>DIS(%)</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;font-size:11px;border-bottom:none;"
                    NOWRAP>BAYAR</th
                  >
                </tr>
              </thead>
              <tbody>
                {#each resultinvoice as rec}
                  <tr>
                    <td
                      class="table-fill text-center"
                      class:dark={daylight === false}
                      NOWRAP>{rec.nomor}</td
                    >
                    <td
                      class="table-fill text-center"
                      class:dark={daylight === false}
                      NOWRAP>{rec.permainan}</td
                    >
                    <td
                      class="table-fill text-end"
                      class:dark={daylight === false}
                      NOWRAP>{new Intl.NumberFormat().format(rec.bet)}</td
                    >
                    <td
                      class="table-fill text-end"
                      class:dark={daylight === false}
                      NOWRAP>{rec.kei.toFixed(1)}</td
                    >
                    <td
                      class="table-fill text-end"
                      class:dark={daylight === false}
                      NOWRAP>{rec.diskon.toFixed(1)}</td
                    >
                    <td
                      class="table-fill text-end"
                      class:dark={daylight === false}
                      NOWRAP>{new Intl.NumberFormat().format(rec.bayar)}</td
                    >
                  </tr>
                {/each}
              </tbody>
            </table>
          </slot:template>
        </PanelFull>
      </div>
    </div>
  {/if}
  <div class="clearfix" />
  <style>
    #stream::-webkit-scrollbar {
      width: 0.3em;
    }

    #stream::-webkit-scrollbar-track {
      -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    }

    #stream::-webkit-scrollbar-thumb {
      background-color: #505768;
      outline: 1px solid slategrey;
    }
    .table-dark {
      --bs-table-bg: #212529;
      --bs-table-striped-bg: #1e1e1e;
      --bs-table-striped-color: #fff;
      --bs-table-active-bg: #373b3e;
      --bs-table-active-color: #fff;
      --bs-table-hover-bg: #323539;
      --bs-table-hover-color: #fff;
      color: #fff;
      border-color: #373b3e;
    }
    .table > :not(:first-child) {
      border-top: none;
    }
  </style>
{:else if statuspasaran == "OFFLINE"}
  <div style="margin-bottom:10px;margin-left: -10px;">
    <center>
      <div style="cursor: pointer;">
        {#if agent_home_url !== ""}
          <a href={agent_home_url} title="nuketoto"> Back to Home </a>
        {:else}
          <a
            href="/?token={client_token}&agent={client_company}"
            title="nuketoto"
          >
            Back to Home
          </a>
        {/if}
      </div>
    </center>
  </div>
{/if}
<Loader2 cssstyle={css_loader} />

<style scoped>
  .nav.top-menu > .nav-item > .nav-link.active,
  .nav.top-menu > .nav-item > .nav-link:hover {
    background: #00a86b;
    border-radius: 4px;
    color: #fff;
    cursor: pointer;
  }

  .nav.top-menu > .nav-item > .nav-link {
    color: #171717;
  }
  .nav.top-menu > .nav-item {
    font-size: 10pt;
    font-family: "Poppins";
    opacity: 1;
    color: #171717;
    width: 105px;
    text-align: center;
    margin: 0 5px;
  }

  .nav.top-menu.dark > .nav-item > .nav-link.active,
  .nav.top-menu.dark > .nav-item > .nav-link:hover {
    color: #171717;
  }

  .nav.top-menu.dark .nav-link {
    color: #fff;
  }

  .nav.top-menu > .nav-item > .nav-link,
  .nav.top-menu > .nav-item > .nav-link.active {
    text-decoration: none;
  }
</style>

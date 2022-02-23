<script lang="ts">
  import { Row, Col, Container } from "sveltestrap";
  import Home from "./pages/Home.svelte";
  import Permainan from "./pages/Permainan.svelte";
  import Loader from "./components/Loader.svelte";
  import Notif from "./components/Notif.svelte";
  import dayjs from "dayjs";
  import utc from "dayjs/plugin/utc";
  import timezone from "dayjs/plugin/timezone";
  import Notification, { notifications } from "./components/Noti.svelte";
  import Swal from "sweetalert2";

  dayjs.extend(utc);
  dayjs.extend(timezone);

  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  const token_browser = urlParams.get("token");
  const agentCode = urlParams.get("agent");
  const marketCode = urlParams.get("market");
  const agent_home = urlParams.get("homeUrl");
  let client_device = "";
  if (token_browser === null) {
    console.log("TOKEN NOT FOUND");
  } else {
    initTimezone();
  }
  if (
    /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
      navigator.userAgent
    )
  ) {
    // true for mobile device
    client_device = "MOBILE";
  } else {
    client_device = "WEBSITE";
  }
  let listkeluaran = [];
  let pasaran_name = "";
  let pasaran_code = "";
  let pasaran_periode = 0;
  let permainan = "4-3-2";
  let client_token = "";
  let client_company = "";
  let client_username = "";
  let client_ipaddress = "";
  let client_timezone = "";
  let client_website_status = "";
  let client_website_message = "";
  let agent_home_url = "";
  $: if (marketCode) {
    pasaran_code = marketCode;
  }

  $: if (agent_home) {
    agent_home_url = agent_home;
  }

  const pasaran = (e) => {
    pasaran_code = e.detail.code;
    pasaran_name = e.detail.name;
    pasaran_periode = e.detail.periode;
  };

  let record = "";
  let message_err = "";
  let css_err = "display:none;";
  let checked = false;
  let revisi_note = "";
  let balance_credit = 0;
  // let daylight = false;
  $: daylight = checked;
  $: client_credit = balance_credit;

  // $: callday = daylights();
  function daylights() {
    const time = new Date().getHours();
    if (time >= 6 && time < 18) {
      return true;
    }
    return false;
  }

  async function initTimezone() {
    const res = await fetch("/api/healthz");
    if (!res.ok) {
      const message = `An error has occured: ${res.status}`;
      throw new Error(message);
    } else {
      const json = await res.json();
      client_ipaddress = json.real_ip;
      client_timezone = "Asia/Jakarta";
    }
    let callday = daylights();
    if (callday) {
      checked = true;
    }
    initapp(token_browser, agentCode);
  }
  async function initapp(token, agent_code) {
    const resInit = await fetch("/api/init", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-forwarder-for": client_ipaddress,
      },
      body: JSON.stringify({
        token,
        agent_code,
      }),
    });
    if (!resInit.ok) {
      const initMessage = `An error has occured: ${resInit.status}`;
      throw new Error(initMessage);
    } else {
      const initJson = await resInit.json();

      if (initJson.status === 200) {
        switch (initJson.company) {
          case "":
            css_err = "display:inline-block";
            message_err = "Agent not found, Please contact admin";
            Swal.fire({
              icon: "error",
              title: "Oops...",
              text: message_err,
              heightAuto: false,
            });
            // notifications.push(message_err, "white", "middle");
            setTimeout(function () {
              css_err = "display: none;";
            }, 5000);
            break;
          default:
            client_token = initJson.token;
            client_company = initJson.company;
            client_username = initJson.developer;
            balance_credit = initJson.credit;
            client_website_status = initJson.website_status;
            client_website_message = initJson.website_message;
            if (client_website_status == "OFFLINE") {
              client_token = "";
              // message_err = client_website_message;
              // css_err = "display:inline-block";
              Swal.fire({
                icon: "warning",
                title: "Sorry we are offline now",
                html: client_website_message.replace(",", "<br/>"),
                heightAuto: false,
              });
            } else {
              initPasaran();
            }
            break;
        }
      }
    }
  }
  async function initPasaran() {
    const resPasar = await fetch("/api/listpasaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        token: client_token,
        company: client_company,
        timezone: client_timezone,
      }),
    });
    if (!resPasar.ok) {
      const pasarMessage = `An error has occured: ${resPasar.status}`;
      throw new Error(pasarMessage);
    } else {
      const jsonPasar = await resPasar.json();
      if (jsonPasar.status == 200) {
        record = jsonPasar.record;
        if (record != null) {
          for (var i = 0; i < record.length; i++) {
            listkeluaran = [
              ...listkeluaran,
              {
                id: record[i]["pasaran_id"],
                pasaran_code: record[i]["pasaran_id"],
                pasaran: record[i]["pasaran_togel"],
                pasaran_periode: record[i]["pasaran_periode"],
                pasaran_tgl: dayjs(record[i]["pasaran_marketclose"])
                  .tz(client_timezone)
                  .format("DD MMM YYYY | HH:mm:ss"),
                pasaran_status: record[i]["pasaran_status"],
              },
            ];
          }
        } else {
          Swal.fire({
            icon: "error",
            title: "Oops...",
            text: "An error has occured, Please Contact Administrator",
            heightAuto: false,
          });
        }
      } else {
        Swal.fire({
          icon: "error",
          title: "Oops...",
          text: "An error has occured, Please Contact Administrator",
          heightAuto: false,
        });
      }
    }
  }
</script>

<svelte:head>
  <title>SDSB4D</title>
</svelte:head>
<Notification duration="3000" />
{#if client_device == "WEBSITE"}
  <div class="content" class:dark={daylight === false}>
    {#if client_token != ""}
      <Container>
        {#if token_browser != ""}
          {#if pasaran_code != ""}
            <Row align-items-start>
              <Permainan
                {client_token}
                {client_company}
                {client_username}
                {client_credit}
                {client_ipaddress}
                {client_timezone}
                {client_device}
                {pasaran_code}
                {pasaran_name}
                {pasaran_periode}
                {permainan}
                bind:checked
                bind:balance_credit
                bind:revisi_note
                {daylight}
                {agent_home_url}
              />
            </Row>
          {:else}
            <Row align-items-start>
              <Home
                {client_token}
                {client_company}
                {client_username}
                {client_credit}
                {client_ipaddress}
                {client_timezone}
                {client_device}
                {listkeluaran}
                bind:checked
                {daylight}
                on:pasaran={pasaran}
              />
            </Row>
          {/if}
        {/if}
      </Container>
    {:else}
      <div style="padding-top: 100px">
        <center>
          <Loader />
        </center>
      </div>
    {/if}
  </div>
{:else}
  <div class="content mobile" class:dark={daylight === false}>
    {#if client_token != ""}
      <Container>
        {#if token_browser != ""}
          <Row align-items-start>
            {#if pasaran_code != ""}
              <Permainan
                {client_token}
                {client_company}
                {client_username}
                {client_credit}
                {client_ipaddress}
                {client_timezone}
                {client_device}
                {pasaran_code}
                {pasaran_name}
                {pasaran_periode}
                {permainan}
                bind:checked
                bind:balance_credit
                {daylight}
              />
            {:else}
              <Home
                {client_token}
                {client_company}
                {client_username}
                {client_credit}
                {client_ipaddress}
                {client_timezone}
                {client_device}
                {listkeluaran}
                bind:checked
                {daylight}
                on:pasaran={pasaran}
              />
            {/if}
          </Row>
        {/if}
      </Container>
    {:else}
      <div style="height: 100%;margin:100px 0px 100px 0px;">
        <center>
          <Loader cssstyle={"height: 100%;margin:100px 0px 100px 0px;"} />
        </center>
      </div>
    {/if}
  </div>
{/if}

<style>
  .content {
    background-image: url("/bg-light.svg");
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center center;
    background-attachment: scroll;
    min-height: 100%;
  }

  .content.dark {
    background-image: url("/bg-dark.svg");
  }

  .content.mobile,
  .content.dark.mobile {
    height: unset;
  }
</style>

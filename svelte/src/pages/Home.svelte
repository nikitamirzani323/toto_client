<script>
  import { createEventDispatcher } from "svelte";
  import {
    Card,
    CardBody,
    CardHeader,
    CardFooter,
    Col,
    Row,
  } from "sveltestrap";
  import Header from "../components/Header.svelte";
  import { notifications } from "../components/Noti.svelte";

  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_credit = 0;
  export let client_ipaddress = "";
  export let client_timezone = "Asia/Jakarta";
  export let client_device = "";
  export let listkeluaran = [];
  export let daylight = false;
  export let checked;
  let home = true;
  let dispatch = createEventDispatcher();

  const handleClick = (code, name, periode, status) => {
    if (status == "ONLINE") {
      const pasaran = {
        code,
        periode,
        name,
        client_token,
      };
      dispatch("pasaran", pasaran);
    } else {
      notifications.push("PASARAN " + name + " OFFLINE");
    }
  };
</script>

<Header
  {client_username}
  {client_company}
  {client_credit}
  {client_token}
  {client_ipaddress}
  {client_timezone}
  {client_device}
  bind:checked
  {daylight}
  {home}
/>
<div class="row">
  {#each listkeluaran as { pasaran_code, pasaran, pasaran_periode, pasaran_status, pasaran_tgl }, i}
    {#if client_device == "WEBSITE"}
      <Col
        xxl="3"
        xl="3"
        lg="6"
        md="6"
        sm="6"
        xs="6"
        style="cursor:pointer;padding:5px;margin:0px;"
      >
        <div
          class="default-block web-block card"
          class:dark={daylight === false}
          on:click={() => {
            handleClick(pasaran_code, pasaran, pasaran_periode, pasaran_status);
          }}
        >
          <div class="card-header" class:custom-dark={daylight === false}>
            <center id="style_text">
              {#if pasaran_status == "ONLINE"}
                <span class="badge rounded-pill bg-online blink_me"
                  >{pasaran_status}</span
                >
              {:else}
                <span class="badge rounded-pill bg-secondary"
                  >{pasaran_status}</span
                >
              {/if}
            </center>
          </div>
          <CardBody>
            <center>
              <img class="mb-4" src="/flags/{pasaran_code}.svg" />
              <br />
              <h4 class="head-fonts" class:custom-dark={daylight === false}>
                {pasaran}
              </h4>
              <span style="font-size: 12px;line-height: 18px;"
                >{pasaran_tgl} WIB</span
              >
              <br />
              <span style="font-size: 10px;line-height: 15px;"
                >PERIODE : {pasaran_periode}</span
              >
            </center>
          </CardBody>
        </div>
      </Col>

      {#if (i + 1) % 4 === 0}
        <div class="clearfix" />
      {/if}
    {:else}
      <Col
        xxl="6"
        xl="6"
        lg="6"
        md="6"
        sm="6"
        xs="6"
        style="cursor:pointer;padding:10px;margin:0px;"
      >
        <div
          class="default-block mobile-block card"
          style="border:none;"
          on:click={() => {
            handleClick(pasaran_code, pasaran, pasaran_periode, pasaran_status);
          }}
        >
          <div class="card-header">
            {#if pasaran_status == "ONLINE"}
              <div
                class="blink_me"
                style="margin-top:0px;color:black;background:#c4f750;font-size:12px;font-weight:bold;padding:5px;border-radius:15px;text-align:center;"
              >
                {pasaran_status}
              </div>
            {:else}
              <div
                style="margin-top:0px;background:#f7602e;color:white;font-size:12px;font-weight:bold;padding:5px;border-radius:15px;text-align:center;"
              >
                {pasaran_status}
              </div>
            {/if}
          </div>
          <CardBody>
            <center id="style_text" style="font-size:15px;">
              <h5 class="head-fonts">{pasaran}</h5>
              <span style="font-size: 11px;">
                PERIODE : {pasaran_periode}
              </span>
              <br />
              <span style="font-size: 10px;">{pasaran_tgl} WIB</span>
            </center>
          </CardBody>
        </div>
      </Col>
    {/if}
  {/each}
</div>

<div class="clearfix" />
<br />

<style>
  .btn-play {
    color: #fff;
    border-color: #ff9900;
    border-radius: 15px;
    margin-top: 3rem;
    padding: 5px 25px;
  }

  .btn.btn-play:hover {
    background: linear-gradient(100.05deg, #ff9900 24.87%, #e56b00 85.21%);
  }
  .bg-online {
    background-color: #009e42 !important;
  }
  .card:hover {
    border: 2px solid #ff9900;
  }
  .card-header {
    border-bottom: none !important;
    padding: 2rem 1rem 0.5rem !important;
    background-color: #fff;
    border-radius: 20px;
  }

  .card-header.custom-dark {
    background-color: #171717;
  }
  .default-block {
    /* background-image: url("/bg-button.svg"),
      linear-gradient(180deg, #8e000e 0%, #5f0009 72.69%); */
    background-color: #fff;
    background-size: cover;
    margin-bottom: 1.5rem;
    margin-top: 1.5rem;
    filter: drop-shadow(0px 4px 20px rgba(0, 0, 0, 0.15));
  }

  .default-block.dark {
    background-color: #171717;
  }

  .web-block {
    width: 227px;
    height: 227px;
    border-radius: 20px;
  }

  #style_text {
    font-size: 0.8rem;
    letter-spacing: 1px;
    font-family: Montserrat;
  }
  @media (min-width: 1200px) {
    .web-block {
      width: 250px;
      height: 250px;
    }
  }

  @media (min-width: 1400px) {
    .web-block {
      width: 300px;
      height: 300px;
    }

    .head-fonts {
      font-size: 2rem;
      color: #171717;
    }

    .head-fonts.custom-dark {
      color: #fff;
    }
    #style_text {
      font-size: 1.5rem;
    }
  }

  .mobile-block {
    width: 100%;
    height: 163px;
    border-radius: 14px;
    margin-bottom: 0;
  }
  .blink_me {
    animation: blinker 1s linear infinite;
  }

  @keyframes blinker {
    50% {
      opacity: 0;
    }
  }
</style>

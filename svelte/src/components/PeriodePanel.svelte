<script>
  import { CardHeader } from "sveltestrap";
  export let pasaran_name;
  export let permainan_title;
  export let pasaran_periode;
  export let pasaran_code;
  export let daylight = false;
</script>

<CardHeader class="card-periode {daylight ? '' : 'dark'}">
  <div class="row">
    <div class="col">
      <h2 class="pt-2">
        {pasaran_name} - {permainan_title}
      </h2>
    </div>
    <div class="col">
      <div class="float-end">
        <h5 class="pt-3 fw-normal text-black-60">
          PERIODE : {pasaran_periode + " - " + pasaran_code}
        </h5>
      </div>
    </div>
  </div>
</CardHeader>

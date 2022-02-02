<script>
  export let count_line;
  export let count_line_4d;
  export let count_line_3d;
  export let count_line_3dd;
  export let count_line_2d;
  export let count_line_2dd;
  export let count_line_2dt;
  export let daylight;
  export let line4d = false;
</script>

<div class="container">
  <div class="row">
    <div class="col-md fs-4">
      TOTAL LINE : <span class="color-font" class:dark={daylight === false}
        >{count_line}</span
      >
    </div>
  </div>
  {#if line4d}
    <div class="row">
      <div class="col fs-6">
        4D : <span class="color-font" class:dark={daylight === false}
          >{count_line_4d}</span
        >
      </div>
      <div class="col fs-6">
        3D : <span class="color-font" class:dark={daylight === false}
          >{count_line_3d}</span
        >
      </div>
      <div class="col fs-6">
        3DD : <span class="color-font" class:dark={daylight === false}
          >{count_line_3dd}</span
        >
      </div>
      <div class="col fs-6">
        2D : <span class="color-font" class:dark={daylight === false}
          >{count_line_2d}</span
        >
      </div>
      <div class="col fs-6">
        2DD : <span class="color-font" class:dark={daylight === false}
          >{count_line_2dd}</span
        >
      </div>
      <div class="col fs-6">
        2DT : <span class="color-font" class:dark={daylight === false}
          >{count_line_2dt}</span
        >
      </div>
    </div>
  {/if}
</div>

<script>
  export let modal_id = "";
  export let modal_size = "";
  export let modal_body_height = "height:350px;overflow:scroll;";
  export let modal_headerbootom_flag = false;
  export let modal_footer_flag = true;
  export let daylight = false;
</script>

<div class="modal fade" id={modal_id} tabindex="-1" data-bs-backdrop="false">
  <div class="modal-dialog {modal_size}">
    <div class="modal-content custom" class:dark={daylight === false}>
      <div class="modal-header custom" class:dark={daylight === false}>
        <slot name="header" />
        <button
          type="button"
          class="btn-close custom"
          class:btn-close-white={daylight === false}
          data-bs-dismiss="modal"
          aria-label="Close"
        />
      </div>
      {#if modal_headerbootom_flag}
        <div style="background-color: transparent">
          <slot name="headerbottom" />
        </div>
      {/if}
      <div
        data-simplebar
        class="modal-body"
        id="streaminvoice"
        style="font-size:12px;padding:0px;margin:0px;background:{daylight
          ? '#fff'
          : '#121212'};border:1px solid {daylight
          ? '#cbcbcb'
          : '#0e0c13'};color:{daylight
          ? '#171717'
          : 'white'};{modal_body_height};
          border-bottom-left-radius:10px;
          border-bottom-right-radius:10px;"
      >
        <slot name="body" />
      </div>
      {#if modal_footer_flag}
        <div
          class="modal-footer"
          style="margin:0px;background: {daylight
            ? '#f3f3f3'
            : '#101010'};border-top: none;
          border-bottom-left-radius:10px;
          border-bottom-right-radius:10px;"
        >
          <slot name="footer" />
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .modal-content.custom {
    background: #fff;
    color: #fff;
    border-radius: 20px;
  }

  .modal-content.custom.dark {
    background: #171717;
    color: #171717;
  }

  .btn-close.custom {
    background: #fff url("/btn-close.svg") center/1em auto no-repeat;
  }
  .modal-header.custom {
    background: linear-gradient(180deg, #00a86b 21.53%, #004029 87.36%);
    border-radius: 20px 20px 0 0;
    padding: 30px 20px;
    border-bottom: none;
    color: #fff;
  }
  .modal-header.custom.dark {
    color: #fff;
  }
  .modal-title.dark {
    color: #171717;
  }
  #streaminvoice::-webkit-scrollbar {
    width: 0.3em;
  }

  #streaminvoice::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(66, 66, 66, 0.3);
  }

  #streaminvoice::-webkit-scrollbar-thumb {
    background-color: #424242;
    outline: 1px solid #424242;
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
</style>

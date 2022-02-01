<script>
  export let checked = false;
  export let color = "#2196F3";
</script>

<label class="switch float-end">
  <input type="checkbox" bind:checked />
  <span class="slider" />
</label>

<style>
  .switch {
    position: relative;
    display: inline-block;
    width: 50px;
    height: 23px;
  }

  .switch input {
    opacity: 0;
    width: 0;
    height: 0;
  }

  .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #ccc;
    -webkit-transition: 0.4s;
    transition: 0.4s;
    border-radius: 20px;
  }

  .slider:before {
    position: absolute;
    content: "";
    height: 18px;
    width: 18px;
    left: 3px;
    top: 3px;
    background-color: white;
    -webkit-transition: 0.4s;
    transition: 0.4s;
    border-radius: 50%;
  }

  input:checked + .slider {
    background-color: #6f7072;
  }

  input:checked + .slider {
    box-shadow: 0 0 1px #6f7072;
  }

  input:checked + .slider:before {
    -webkit-transform: translateX(26px);
    -ms-transform: translateX(26px);
    transform: translateX(26px);
  }
</style>

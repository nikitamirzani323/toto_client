<script>
    export let modal_id = "";
    export let modal_size = "";
</script>

<div
    class="modal fade"
    id={modal_id}
    tabindex="-1"
    data-bs-backdrop="false">
    <div class="modal-dialog {modal_size}">
        <div class="modal-content">
            <div
                data-simplebar
                class="modal-body"
                id="streaminvoice"
                style="font-size:12px;padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;color:white;">
                <slot name="body" />
            </div>
        </div>
    </div>
</div>

<script>
  import { Card, CardBody, CardHeader } from "sveltestrap";
  export let daylight = false;
</script>

<Card class="panel-menus {daylight ? '' : 'dark'}">
  <CardBody style="padding:10px;margin:0px;">
    <slot name="body" />
  </CardBody>
</Card>

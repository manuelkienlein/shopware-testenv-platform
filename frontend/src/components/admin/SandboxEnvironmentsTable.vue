<script>
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from "primevue/button";
import Tag from "primevue/tag"
import SandboxService from "../../services/sandboxService.js"

export default {

  components: {
    DataTable,
    Column,
    Button,
    Tag
  },

  // Properties returned from data() become reactive state
  // and will be exposed on `this`.
  data() {
    return {
      sandboxes: [{
        id: "aafb929f-0992-44d2-a8d3-17619360deff",
        container_id: "1fdab6cd2c18d260586b462ec7cd86f482d5b7fef6ee84fdb3733b97b79ae652",
        container_name: "/sandbox-aafb929f-0992-44d2-a8d3-17619360deff",
        url: "sandbox-aafb929f-0992-44d2-a8d3-17619360deff.shopshredder.zion.mr-pixel.de",
        image: "dockware/dev:6.6.8.2",
        created_at: "2024-12-20T13:08:50+01:00",
        state: "running",
        status: "Up 1 second"
      }]
    }
  },

  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event handlers in templates.
  methods: {
    async loadData() {
      this.sandboxes = await SandboxService.getAllSandboxes();
    },
    getStatus(sandbox) {
      switch (sandbox.state) {
        case 'running':
          return 'success';
        default:
          return "danger";
      }
    },
    openSandboxWindow(data) {
      window.open(data.url, '_blank').focus();
    },
    async deleteSandbox(data) {
      await SandboxService.deleteSandbox(data.id);
      await this.loadData();
    }
  },

  // Lifecycle hooks are called at different stages
  // of a component's lifecycle.
  // This function will be called when the component is mounted.
  mounted() {
    console.log(`Loading sandbox table`)
    this.loadData();
  }
}
</script>

<template>
  <div class="card">
    <DataTable :value="sandboxes" tableStyle="min-width: 50rem">
      <template #header>
        <div class="flex flex-wrap items-center justify-between gap-2">
          <span class="text-xl font-bold">Sandbox Environments</span>
          <Button icon="pi pi-refresh" rounded raised @click="loadData"/>
        </div>
      </template>
      <Column field="id" header="ID"></Column>
      <Column field="image" header="Image"></Column>
      <Column field="created_at" header="Created"></Column>
      <Column field="state" header="Status">
        <template #body="slotProps">
          <Tag :value="slotProps.data.state" :severity="getStatus(slotProps.data)" />
        </template>
      </Column>
      <Column class="w-24 !text-end">
        <template #body="{ data }">
          <div class="flex /*flex-wrap*/ gap-1 justify-center">
            <Button icon="pi pi-arrow-right" @click="openSandboxWindow(data)" severity="secondary" rounded></Button>
            <Button icon="pi pi-trash" @click="deleteSandbox(data)" severity="secondary" rounded></Button>
          </div>
        </template>
      </Column>
      <template #empty>
        <div class="text-center text-gray-500">
          <i class="pi pi-info-circle text-xl"></i>
          <p>No data available!</p>
        </div>
      </template>
    </DataTable>
  </div>
</template>

<style scoped>

</style>
<script>
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from "primevue/button";
import Tag from "primevue/tag"
import SandboxService from "../../services/sandboxService.js";
import ImagesService from "../../services/imagesService.js";

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
      images: [{
        "id": "a407dee395ed97ead1e40c7537395d6271c07cc89c317f8eda1c19f6fc783695",
        "image_name": "dockware/dev",
        "image_tag": "6.6.8.2",
        "created_at": "2024-11-12T17:10:49+01:00",
        "size": 4860039980
      }]
    }
  },

  // Methods are functions that mutate state and trigger updates.
  // They can be bound as event handlers in templates.
  methods: {
    async loadData() {
      this.images = await ImagesService.getAllImages();
    },
    formatSize(bytes) {
      const sizes = ["Bytes", "KB", "MB", "GB", "TB"];
      if (bytes === 0) return "0 Bytes";
      const i = Math.floor(Math.log(bytes) / Math.log(1024));
      const size = bytes / Math.pow(1024, i);
      return `${size.toFixed(2)} ${sizes[i]}`;
    }
  },

  // Lifecycle hooks are called at different stages
  // of a component's lifecycle.
  // This function will be called when the component is mounted.
  mounted() {
    console.log(`Loading images table`)
    this.loadData();
  }
}
</script>

<template>
  <div class="card">
    <DataTable :value="this.images" tableStyle="min-width: 50rem">
      <template #header>
        <div class="flex flex-wrap items-center justify-between gap-2">
          <span class="text-xl font-bold">Sandbox Images</span>
          <Button icon="pi pi-refresh" rounded raised @click="loadData"/>
        </div>
      </template>
      <Column field="id" header="ID">
        <template #body="{ data }">
          {{ data.id.slice(0, 10) }}...
        </template>
      </Column>
      <Column field="image_name" header="Image Name"></Column>
      <Column field="image_tag" header="Image Tag"></Column>
      <Column field="created_at" header="Created"></Column>
      <Column field="size" header="Size">
        <template #body="{ data }">
          {{ formatSize(data.size) }}
        </template>
      </Column>
      <Column class="w-24 !text-end">
        <template #body="{ data }">
          <div class="flex /*flex-wrap*/ gap-1 justify-center">
            <Button icon="pi pi-trash" severity="secondary" rounded></Button>
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
<script>
import Layout from '@layouts/main.vue'

export default {
  page: {
    title: 'Import',
    meta: [{ name: 'Import from other apps', content: 'The Import page.' }],
  },
  components: { Layout },
  props: {
    user: {
      type: Object,
      required: true,
    },
  },
}
</script>

<template>
  <Layout>
    <div class="columns box">
      <div class="column">
        <h1 class="title">{{ $t('importdata') }}</h1>
        <p class="subtitle">{{ $t('importdatadesc') }}</p>
      </div>
    </div>
    <br />
    <div class="columns">
      <div class="column is-one-third">
        <div class="box">
          <h1 class="title">Fuelly</h1>
          <p>If you have been using Fuelly to store your vehicle data, export the CSV file from Fuelly and click here to
            import.</p>
          <br />
          <b-button type="is-primary" tag="router-link" to="/import/fuelly">{{ $t('import') }}</b-button>
        </div>
      </div>

      <div class="column is-one-third" to="/import-fuelly">
        <div class="box">
          <h1 class="title">Drivvo</h1>
          <p>{{ $t('importcsv', { 'name': 'Fuelly' }) }}</p>
          <br />
          <b-button type="is-primary" tag="router-link" to="/import/drivvo">{{ $t('import') }}</b-button>
        </div>
      </div>


      <div class="column is-one-third" to="/import-generic">
        <div class="box">
          <h1 class="title">{{ $t('importgeneric') }}</h1>
          <p>{{ $t('importgenericdesc') }}</p>
          <br />
          <b-button type="is-primary" tag="router-link" to="/import/generic">{{ $t('import') }}</b-button>
        </div>
      </div>
    </div>
  </Layout>
</template>

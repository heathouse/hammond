<script>
import Layout from '@layouts/main.vue'

export default {
  page() {
    return {
      title: this.user.name,
      meta: [
        {
          name: 'description',
          content: `The user profile for ${this.user.name}.`,
        },
      ],
    }
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
    <h1>
      <BaseIcon name="user" />
      {{ user.name }}
      {{ $t('profile') }}
    </h1>
    <pre>{{ user }}</pre>
  </Layout>
</template>

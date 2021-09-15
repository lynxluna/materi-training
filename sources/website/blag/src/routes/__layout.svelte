<script>
  import "../app.postcss";
  import Navigation from "$lib/components/Navigation.svelte";
  import Footer from "$lib/components/Footer.svelte";
</script>

<Navigation />

<slot></slot>

<Footer />


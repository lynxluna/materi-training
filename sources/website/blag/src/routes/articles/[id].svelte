<svelte:head>
  <title>{title}</title>
</svelte:head>
<main>
<section>
  <h1>{title}</h1>
</section>
<section>
  <div>{content}</div>
</section>
</main>
<script>
  import {page} from '$app/stores'
  import { onMount } from 'svelte';
  import { API } from '$lib/api';
  
  let title = "Fetching ..."
  let content = "Getting content ..."
  let articleID = $page.params.id
  
  onMount(async() => {
    const response = await API.single(articleID)
    const resp = await response.json()

    if (response.ok) {
      title = resp.title 
      content = resp.content
    } else {
      window.alert(resp.message)
    }
  })
</script>

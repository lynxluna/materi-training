<main>
  <div class="flex flex-col">
    <h1>New Article</h1>
    <form class="flex flex-col form-control">
      <label for="title" class="pt-4 py-2 label">Title</label>
      <input type="text" name="title" id="title" class="input input-bordered input-primary" bind:value={title}>

        <label for="content" class="pt-4 py-2 label">Content</label>
        <textarea class="h-24 resize-y textarea textarea-primary textarea-bordered" bind:value={content}></textarea>
        <button type="submit" class="btn btn-wide btn-m btn-primary mt-4 mx-auto" on:click|preventDefault={submit}>
            Create
        </button>
    </form>
  </div>
</main>

<svelte:head>
  <title>Create Article</title>
</svelte:head>

<script>
import { goto } from "$app/navigation";

  import {API} from "$lib/api"

  let title = null
  let content = null

  async function submit() {
    const response = await API.create(title, content)
    const resp = await response.json()

    if (response.ok) {
      const aid = resp.id
      goto('/articles/' + aid)
    } else {
      window.alert(resp.message)
    }
  }
</script>

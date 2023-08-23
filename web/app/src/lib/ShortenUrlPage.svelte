<script lang="ts">
  import {SvelteToast, toast} from '@zerodevx/svelte-toast'
  import ShortenedUrl from "./shortUrlDisplay.svelte";

  interface Response {
    key: string;
    long_url: string;
    short_url: string;
  }

  let url = '';
  let data: Response;

  async function shortenUrl() {
    const response = await fetch('http://localhost:8080/app', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({url}),
    });

    if (response.ok) {
      data = await response.json();
    } else {
      toast.push('Invalid URL', {classes: ['warn'], duration: 1000},)
    }
  }

</script>

<SvelteToast options={{ classes: ['log'] }}/>

<div>
  {#if !data}
    <form>
      <div>
        <div class="text-3xl font-semibold m-4">Shorten a long URL</div>
        <div>
          <p class="text-gray-600 font-semibold">Paste a long URL</p>
          <input
            class="w-full text-lg px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-300"
            type="text"
            bind:value={url}/>
        </div>
        <button
          class="m-5 bg-orange-400 hover:bg-orange-800 text-white px-4 py-2 rounded-lg transition duration-300 ease-in-out"
          on:click|preventDefault={shortenUrl}>
          Shorten URL
        </button>
      </div>
    </form>
  {:else}
    <ShortenedUrl {data}/>
  {/if}
</div>

<style>
  :global(.log.info) {
    --toastBackground: #3a943a;
  }

  :global(.log.warn) {
    --toastBackground: #f35959;
  }
</style>

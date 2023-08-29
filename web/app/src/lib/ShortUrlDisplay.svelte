<script lang="ts">
  import {toast} from "@zerodevx/svelte-toast";
  import type {Resp} from "./ShortenUrlPage.types.svelte";

  export let data: Resp;

  function copyToClipboard() {
    const textField = document.createElement('textarea');
    textField.value = data.short_url;
    document.body.appendChild(textField);
    textField.select();
    document.execCommand('copy');
    document.body.removeChild(textField);
    toast.push('Link copied', {classes: ['info'], duration: 1000},)
  }

  function visitUrl() {
    window.location.href = data.short_url;
  }

  function goToMainPage() {
    window.location.href = '/';
  }

</script>

<div>
  <div class="text-3xl font-semibold m-4">Shorten a long URL</div>
  <div>
    <p class="text-gray-600 font-semibold mt5">Long URL</p>
    <p
      class="w-full text-lg px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-300 bg-white">{data.long_url}</p>
  </div>
  <div>
    <p class="text-gray-600 font-semibold mt-5">Short URL</p>
    <div class="flex">
      <p
        class="w-full text-lg px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-300 bg-white">{data.short_url}</p>
      <button id="copy" on:click|preventDefault={copyToClipboard}>Copy</button>
    </div>
  </div>
  <div class="socials mt-5">
    <button on:click|preventDefault={visitUrl}>Visit</button>
    <button on:click={goToMainPage}>Shorten Another</button>
  </div>
</div>
<script lang="ts">
    import type { Nweet } from "$lib/types/nweet.type";
    export let Data: Nweet;
</script>

<div class="card">
    <p class="content">{Data.Content}</p>
</div>

<style>
    .card {
        display: block;
        line-height: 1.6rem;
    }
</style>

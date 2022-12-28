<script lang="ts">
    import PostCard from "$lib/PostCard.svelte";
    import Separator from "$lib/Separator.svelte";
    import { onMount } from "svelte";
    import type { Nweet } from "$lib/types/nweet.type";

    let Posts: Nweet[] = [];
    onMount(async () => {
        const data = await fetch('/nweets', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
		});
        Posts = await data.json();
        console.log(Posts);
    });
</script>


<svelte:head>
    <title>Nweets</title>
</svelte:head>


{#each Posts as post}
    <PostCard Data={post} />
    <Separator />
{/each}

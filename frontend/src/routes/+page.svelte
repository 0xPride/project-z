<script lang="ts">
    import Container from "$lib/Container.svelte";
    import PostCard from "$lib/PostCard.svelte";
    import Separator from "$lib/Separator.svelte";
    import Header from "$lib/Header.svelte";
    import { onMount } from "svelte";

    let Posts = [];
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

<Container>
    <Header />
    {#each Posts as post}
        <PostCard Data={post} />
        <Separator />
    {/each}
</Container>

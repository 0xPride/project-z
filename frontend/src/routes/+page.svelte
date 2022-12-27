<script lang="ts">
    import Container from "$lib/Container.svelte";
    import PostCard from "$lib/PostCard.svelte";
    import Separator from "$lib/Separator.svelte";
    import Header from "$lib/Header.svelte";
    import posts from "./posts.json";
    import TypeIt from "typeit";
    import { onMount } from "svelte";

    onMount(async () => {
        new TypeIt(".typing", {
            strings: "Talk to be heard ......",
            speed: 40,
            waitUntilVisible: false,
        }).go();
    });
</script>


<svelte:head>
    <title>Nweets</title>
</svelte:head>

<Container>
    <Header />
    <div class="typing"></div>
    {#each posts as post}
        <PostCard Data={post} />
        <Separator />
    {/each}
</Container>

<style>
    .typing {
        font-size: 1.5rem;
        font-style: 200;
        padding: 20px;
        margin-bottom: 50px;
        text-align: center;
    }
</style>

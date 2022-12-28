<script>
    import { page } from '$app/stores';
    import Container from '$lib/Container.svelte';
    import Header from '$lib/Header.svelte';
</script>
  
<h1>{$page.error?.message}</h1>
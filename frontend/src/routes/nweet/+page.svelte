<script lang="ts">
    import { goto } from "$app/navigation";

    let messageContent: string = '';
    $: status = `${512 - messageContent.length > 0 ? 512 - messageContent.length + ' charcter left' : 'Message cannot be more than 512 character'}`
    let submitButton : HTMLButtonElement;
    async function handleSubmit() {

        if (messageContent.length > 512 || messageContent.length == 0) {
            if (messageContent.length == 0)
                inform('Message cannot be empty')
            else
                inform('Message too short')
            return;
        }

        const req = fetch("/nweets", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                Content: messageContent,
            }),
        });
        submitButton.setAttribute('disabled', 'true');
        inform('Message being sent ...');
        await req;
        goto('/');
    }

    function inform(message: string) {
        status = message;
    }
</script>

<svelte:head>
    <title>Nweetting</title>
</svelte:head>

<div class="note">
    <span class="sec-title">Note</span>
    <span>Please be respectful and responsible in your messages.<br />we trust you!</span>
</div>

<div class="form">
    <span class="sec-title">Message</span>
    <textarea id="content" bind:value={messageContent} />
    <span class="status">{status}</span>
    <button bind:this={submitButton} on:click={handleSubmit}>Publish</button>
</div>

<style>
    .sec-title {
        font-weight: bold;
        margin-bottom: 10px;
        font-size: 1.3rem;
        padding-bottom: 0.5rem;
        width: 100%;
        display: block;
    }

    textarea#content {
        display: block;
        font-family: "Roboto", sans-serif;
        font-size: 1.1rem;
        border: 2px solid black;
        box-sizing: border-box;
        padding: 15px;
        border-radius: 5px;
        height: 500px;
        width: 100%;
    }

    textarea#content:focus {
        outline: none !important;
    }

    button {
        margin-top: 15px;
        color: white;
        background-color: black;
        padding: 15px 0;
        border-radius: 5px;
        border: none;
        font-weight: bold;
        transition: 100ms;
        font-size: 1.1rem;
        width: 100%;
    }

    button:hover {
        cursor: pointer;
        transform: translateY(-2px);
    }

    button:disabled {
        opacity: 0.8;
    }

    .note {
        width: 100%;
        margin-bottom: 1.5rem;
    }
</style>

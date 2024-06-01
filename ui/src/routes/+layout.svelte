<script lang="ts">
	import { onMount } from 'svelte';
	import '../app.css';
	import { userStore } from '../stores/user';
	import { goto } from '$app/navigation';

    let currentUser: any

    userStore.subscribe((user) => {
        currentUser = user;
    });

    onMount(() => {
        if (!currentUser) {
            goto('/login');
        }
    });
</script>

<div>
	<slot />
</div>

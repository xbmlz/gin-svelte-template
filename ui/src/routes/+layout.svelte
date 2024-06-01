<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
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

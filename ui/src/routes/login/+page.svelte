<script lang="ts">
	import { Card } from 'flowbite-svelte';

	export let site = {
		name: 'Gin Svelte Template',
		img: '/favicon.png',
		link: '/',
		imgAlt: 'Gin Svelte Template'
	};
</script>

<main class="bg-gray-50 dark:bg-gray-900 w-full">
	<div
		class="flex flex-col items-center justify-center px-6 pt-8 mx-auto md:h-screen pt:mt-0 dark:bg-gray-900"
	>
		<a
			href={site.link}
			class="flex items-center justify-center mb-8 text-2xl font-semibold lg:mb-10 dark:text-white"
		>
			<img src={site.img} class="mr-4 h-11" alt={site.imgAlt} />
			<span>{site.name}</span>
		</a>

		<!-- Card -->
		<Card class="w-full" size="md" border={false}>
			<h1 class="text-2xl font-bold text-gray-900 dark:text-white">Sign in to platform</h1>
			<form class="mt-8 space-y-6" on:submit|preventDefault>xxx</form>
		</Card>
	</div>
</main>

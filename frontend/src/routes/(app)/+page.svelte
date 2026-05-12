<script lang="ts">
	import * as Card from "$lib/components/ui/card";
	import * as Table from "$lib/components/ui/table";
	import { Badge } from "$lib/components/ui/badge";
	import { Button } from "$lib/components/ui/button";
	import { Video, Activity, Zap, Clock } from "lucide-svelte";
	import AnimatedCounter from "$lib/components/animated-counter.svelte";

	// Data dummy untuk tampilan (Nanti akan di-fetch dari Golang)
	const stats = [
		{ label: "Total Matches", value: 124, icon: Video, color: "text-blue-500" },
		{ label: "Whiff Punishes", value: 452, icon: Zap, color: "text-yellow-500" },
		{ label: "Processing Tasks", value: 3, icon: Activity, color: "text-green-500" },
	];

	const recentMatches = [
		{ id: "M001", game: "Street Fighter 6", char: "Ken vs Juri", status: "Completed", date: "2 jam yang lalu" },
		{ id: "M002", game: "Tekken 7", char: "Kazuya vs Paul", status: "Processing", date: "10 menit yang lalu" },
		{ id: "M003", game: "SF6", char: "Ryu vs Chun-Li", status: "Completed", date: "1 hari yang lalu" },
	];
</script>

<div class="flex flex-col gap-6 p-6">
	<div class="flex items-center justify-between">
		<div>
			<h1 class="text-3xl font-bold tracking-tight">Dashboard hitboxMetrics</h1>
			<p class="text-muted-foreground">Analisis teknis dan performa pertarungan kamu.</p>
		</div>
		<Button href="/upload" class="bg-blue-600 hover:bg-blue-700">
			<Video class="mr-2 h-4 w-4" /> Analyze New Match
		</Button>
	</div>

	<div class="grid gap-4 md:grid-cols-3">
		{#each stats as stat}
			<Card.Root>
				<Card.Header class="flex flex-row items-center justify-between pb-2 space-y-0">
					<Card.Title class="text-sm font-medium">{stat.label}</Card.Title>
					<stat.icon class="h-4 w-4 {stat.color}" />
				</Card.Header>
				<Card.Content>
					<div class="text-2xl font-bold">
						<AnimatedCounter value={stat.value} />
					</div>
				</Card.Content>
			</Card.Root>
		{/each}
	</div>

	<Card.Root>
		<Card.Header>
			<Card.Title>Recent Analysis</Card.Title>
			<Card.Description>Pertandingan terakhir yang diunggah ke sistem.</Card.Description>
		</Card.Header>
		<Card.Content>
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head>ID</Table.Head>
						<Table.Head>Game</Table.Head>
						<Table.Head>Matchup</Table.Head>
						<Table.Head>Status</Table.Head>
						<Table.Head class="text-right">Waktu</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each recentMatches as match}
						<Table.Row>
							<Table.Cell class="font-medium">{match.id}</Table.Cell>
							<Table.Cell>{match.game}</Table.Cell>
							<Table.Cell>{match.char}</Table.Cell>
							<Table.Cell>
								<Badge variant={match.status === 'Completed' ? 'default' : 'secondary'}>
									{match.status}
								</Badge>
							</Table.Cell>
							<Table.Cell class="text-right text-muted-foreground">{match.date}</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		</Card.Content>
	</Card.Root>
</div>
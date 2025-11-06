import { Component } from "@angular/core";
import { NgOptimizedImage } from "@angular/common";
import { RouterLink, RouterLinkActive } from "@angular/router";

@Component({
	selector: "app-site-header",
	imports: [NgOptimizedImage, RouterLink, RouterLinkActive],
	template: `
		<header
			class="flex justify-between p-4 lg:px-12 border-b border-b-accent"
		>
			<div class="flex gap-2 place-items-center">
				<img
					class="size-8"
					[ngSrc]="'favicon.png'"
					width="512"
					height="512"
					alt=""
				/>
				<span class="font-title text-xl">foxCal</span>
			</div>
			<nav>
				<ul class="flex flex-wrap gap-4 text-lg">
					<a
						class="border-b-transparent border-b-2 p-1 flex gap-1 place-items-center"
						routerLink=""
						routerLinkActive="!border-b-primary"
						[routerLinkActiveOptions]="{ exact: true }"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							class="size-5"
						>
							<path
								fill="currentColor"
								fill-rule="evenodd"
								d="M9.293 2.293a1 1 0 0 1 1.414 0l7 7A1 1 0 0 1 17 11h-1v6a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-6H3a1 1 0 0 1-.707-1.707l7-7Z"
								clip-rule="evenodd"
							/>
						</svg>
						Home
					</a>
					<a
						class="border-b-transparent border-b-2 p-1 flex gap-1 place-items-center"
						routerLink="about"
						routerLinkActive="!border-b-primary"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							class="size-5"
						>
							<path
								fill="currentColor"
								fill-rule="evenodd"
								d="M18 10a8 8 0 1 1-16 0 8 8 0 0 1 16 0Zm-7-4a1 1 0 1 1-2 0 1 1 0 0 1 2 0ZM9 9a.75.75 0 0 0 0 1.5h.253a.25.25 0 0 1 .244.304l-.459 2.066A1.75 1.75 0 0 0 10.747 15H11a.75.75 0 0 0 0-1.5h-.253a.25.25 0 0 1-.244-.304l.459-2.066A1.75 1.75 0 0 0 9.253 9H9Z"
								clip-rule="evenodd"
							/>
						</svg>
						About
					</a>
				</ul>
			</nav>
		</header>
	`,
	styles: ``,
})
export class SiteHeader {}

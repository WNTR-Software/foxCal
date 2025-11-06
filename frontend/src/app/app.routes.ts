import { Routes } from "@angular/router";
import { HomeView } from "./views/home/home.view";
import { AboutView } from "./views/about/about.view";

export const routes: Routes = [
	{
		title: "foxCal",
		path: "",
		component: HomeView,
	},
	{
		title: "foxCal - About",
		path: "about",
		component: AboutView,
	}
];

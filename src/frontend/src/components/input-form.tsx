"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import { Switch } from "./ui/switch"
import { Label } from "./ui/label"

const formSchema = z.object({
  StartKeyword: z.string().min(1, { message: "Start keyword must be filled"}).max(50, { message: "Start keyword maximum is 50 characters" }),
  SearchKeyword: z.string().min(1, { message: "Search keyword must be filled"}).max(50, { message: "Search keyword maximum is 50 characters" }),
  MaxIteration: z.number().int().min(1, { message: "Minimum iterasi adalah 1"}).max(100, { message: "Maximum iterasi adalah 100" }).optional(),
})


export function InputForm() {

  const [algorithm, setAlgorithm] = useState<"bfs" | "ids">("ids");


  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      StartKeyword: "",
      SearchKeyword: "",
      MaxIteration: 1,
    },
    mode: "all"
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values)
  }

  // 3. Render your form.
  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <h2 className="text-xl font-semibold text-muted-foreground" >Algorithm : {algorithm === "bfs" ? "Breadth First Search" : "Iterative Deepening Search"}</h2>
        <FormField
          control={form.control}
          name="StartKeyword"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Start Keyword</FormLabel>
              <FormControl>
              <Input placeholder="Kucing" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="SearchKeyword"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Search Keyword</FormLabel>
              <FormControl>
              <Input placeholder="Bahasa Jawa" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="MaxIteration"
          disabled={algorithm === "bfs"}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Maximum Iteration</FormLabel>
              <FormControl>
              <Input placeholder="1" type="number" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex gap-5 items-center">
          <Label className="text-muted-foreground">Breadth First Search</Label>
          {/* <Switch className="" onCheckedChange={setAlgorithm(algorithm === "bfs" ? "ids" : "bfs")}></Switch> */}
          <Label className="text-muted-foreground">Iterative Deepening Search</Label>
        </div>
        <Button type="submit">Start</Button>
      </form>
    </Form>

  )
}
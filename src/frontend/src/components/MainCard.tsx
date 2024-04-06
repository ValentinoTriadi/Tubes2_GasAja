import { InputForm } from "./input-form";
import { Card, CardContent, CardHeader } from "./ui/card";

export const MainCard = () => {
    return (
        <Card className="w-full sm:w-[99%] md:w-[97%] lg:w-[95%] bg-card rounded-md bg-clip-padding backdrop-filter backdrop-blur-md bg-opacity-70 border border-border z-0">
            <CardHeader>
                <h1 className="text-3xl font-bold">Welcome to <strong>GasAja</strong> WikiRace!</h1>
                <hr className="border-t-2 border-primary mt-2 w-full" />
            </CardHeader>
            <CardContent>
                <InputForm></InputForm>
            </CardContent>
        </Card>
    );
};
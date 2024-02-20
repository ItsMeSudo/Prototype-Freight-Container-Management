import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import React from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function ADD_CSV(): JSX.Element {
  const [File, setFile] = React.useState<any>("");
  const [response, updateResponse] = React.useState({ success: 0, incorrectPositions: [] });

  const handleFileChange = (event: any) => {
    setFile(event.target.files);
  };

  const handleSubmit = async () => {
    const data = new FormData();
    data.append("file", File[0]);
    if (!File[0]) return;
      const resp  = await fetch("http://172.16.0.2:3001/api/v2/importcsv", {
        method: "POST",
        body: data
      });

      const jsonResponse = await resp.json();

      if (resp.ok) {
        updateResponse(jsonResponse);
      } else {
        throw new Error('Failed to upload file');
      }
  };

  return (
    <>
      <div>
        <Card>
          <CardHeader>
            <CardTitle>Import from CSV</CardTitle>
            <CardDescription>You can add your CSV</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid w-full max-w-sm items-center gap-1.5">
              <Label>Comma-Separated Values</Label>
              <Input accept=".csv" type="file" onChange={handleFileChange} />
            </div>
          </CardContent>
          <CardFooter className="flex justify-start w-[100%] items-start">
            <Dialog>
              <DialogTrigger onClick={handleSubmit} asChild>
                <Button className="w-[100%]">Add your CSV</Button>
              </DialogTrigger>
              <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                  <DialogTitle>Result of response</DialogTitle>
                  <DialogDescription>
                    You can see here the response of your uploaded file
                  </DialogDescription>
                </DialogHeader>
                <div className="grid gap-4 py-4">
                  {response && (
                    <div className="font-inter-light text-center text-gray-300">
                      Succesfull: {response.success}
                      <div></div>
                      Incorrect: {response.incorrectPositions.length}
                    </div>
                  )}
                </div>
              </DialogContent>
            </Dialog>
          </CardFooter>
        </Card>
      </div>
    </>
  );
}

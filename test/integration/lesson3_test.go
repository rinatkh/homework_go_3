package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCommandOutput(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "errors",
			path: "./cmd/01_errors",
			want: "nil=true required=email: is required age=26 parseErr=false notFound=true field=email count=2 first=email: is required class=not_found",
		},
		{
			name: "arrays",
			path: "./cmd/02_arrays",
			want: "first=10 last=d middle=[1 8 3] swap=[4 2 3 1] sum=9 reverse=[5 4 3 2 1] equal=true original=[4 5 6] changed=[4 99 6] max=-2 diagonal=15 rotate=[4 1 2 3]",
		},
		{
			name: "structs",
			path: "./cmd/03_structs",
			want: "user={ID:7 Name:Maria Active:false} renamed={ID:7 Name:Masha Active:false} active=true sameID=true label=7:Masha product={Name:course Price:2500} discounted={Name:course Price:2000} order={ID:15 Amount:4000 Paid:true} status=paid",
		},
		{
			name: "new_make",
			path: "./cmd/04_new_make",
			want: "int=26 string=go slice=[0 0] len=2 cap=5 nil=false appended=[10 20 30] emptyNil=false user=Maria/3 newSliceNil=true mapNil=false chanCap=2",
		},
		{
			name: "slices",
			path: "./cmd/05_slices",
			want: "first=10/true last=sql/true regular=2/3 limited=2/2 mutated=[10 999 30 40] appendSource=[10 20 30 777] appendPart=[20 30 777] limitedSource=[10 20 30 40] limitedPart=[20 30 777] grow=true kind=filled independent=[1 2 3 4] original=[1 2]",
		},
		{
			name: "loops",
			path: "./cmd/06_loops",
			want: "sum=15 between=18 countdown=[4 3 2 1] factorial=120 even=3 negative=-2/true noZero=3 limited=7 doubled=[2 -4 0 8] replaced=[1 9 0 4] active=2 names=[Maria Ira] runes=3 indexes=[0 2 3] repeated=[1 1 2 2]",
		},
		{
			name: "functions",
			path: "./cmd/07_functions",
			want: "divide=5/false user=Maria/true sum=10 apply=13 allowed=16 counter=6,7 accumulator=15,13 multiply=21 defer=body-second-first captured=first deferred=second named=10 cleanup=[action cleanup] operation=20/true",
		},
		{
			name: "panics",
			path: "./cmd/08_panics",
			want: "normal=false boom=true message=broken parse=42/true invalid=0/false index=20/true missing=0/false error=panic: bad state events=[body defer] outside=false sequence=2/1 pointer=7",
		},
		{
			name: "common",
			path: "./cmd/09_common",
			want: `report="active=2 names=Maria,Ira"/false order={ID:7 Amount:1000 Paid:false} discounted={ID:7 Amount:850 Paid:false}/false source=[10 20 30 40] window=[20 30 777]/false`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", tt.path)
			cmd.Dir = "../.."

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run %s failed: %v\n%s", tt.path, err, output)
			}

			got := strings.TrimSpace(string(output))
			if got != tt.want {
				t.Fatalf("unexpected output for %s:\n\ngot:\n%s\n\nwant:\n%s", tt.path, got, tt.want)
			}
		})
	}
}

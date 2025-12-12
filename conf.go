package main

import (
	"fmt"
	"os/exec"
	"strings"

	// "io/fs"
	"os"
	// "path/filepath"
	// "runtime"
	// "strconv"
	// "strings"

	"github.com/dhamith93/systats"
	// "github.com/shirou/gopsutil/v4/mem"
)

var asciiArt = "    _                       \n    \\`*-.                   \n     )  _`-.                \n    .  : `. .               \n    : _   '  \\              \n    ; " + "*" + "` _.   `*-._         \n    `-.-'          `-.      \n      ;       `       `.    \n      :.       .        \\   \n      . \\  .   :   .-'   .  \n      '  `+.;  ;  '      :  \n      :  '  |    ;       ;-.\n      ; '   : :`-:     _.`* ;\n   .*' /  .*' ; .*`- +'  `*'\n   `*-*   `*-*  `*-*'";

var accentCol string = TERM_BOLD + FORE_BLUE;
var stats = systats.New();

var CONFIG = Config {
  accentCol: accentCol,
  lines: []Line {
    {
      desc: "host",
      prependDesc: false,
      lnFormat: "%s",
      formatSet: func() []any{
        getHost, _ := os.Hostname();
        return []any { accentCol + getHost };
      },
    },

    {
      desc: "div",
      prependDesc: false,
      lnFormat: "%s_________________________________________________",
      formatSet: func() []any { return []any { FORE_BLUE }; },
    },
    
    {
      desc: "wm",
      prependDesc: true,
      lnFormat: "%s",
      formatSet: func() []any{
        // if runtime.GOOS == "darwin" {
        //   return []any { "mac" };
        // }
        getWM := os.Getenv("XDG_CURRENT_DESKTOP");
        return []any { getWM };
      },
    },

    {
      desc: "mem",
      prependDesc: true,
      lnFormat: "%s/%s%s",
      formatSet: func() []any {
        var avm string;
        var tot string;
        // if runtime.GOOS == "darwin" {
        //   v, _ := mem.VirtualMemory();
        //   avm = strconv.Itoa(int(v.Total / 2048));
        //   tot = strconv.Itoa(int(v.Used / 2048));
        // }
        // if runtime.GOOS == "linux" {
          get, _ := stats.GetMemory(systats.Gigabyte);
          getSwap, _ := stats.GetSwap(systats.Gigabyte);
          avm = fmt.Sprint(get.Used/1e+6, get.Unit);
          tot = fmt.Sprint(get.Total/1e+6, TERM_RESET, get.Unit);
          per := fmt.Sprint(int(get.PercentageUsed), "%");
          if int(get.PercentageUsed) < 45 {
            per = FORE_YELLOW + per + TERM_RESET;
            avm = FORE_YELLOW + avm + TERM_RESET;
          }
          if int(get.PercentageUsed) > 45 {
            per = FORE_RED + per + TERM_RESET;
            avm = FORE_RED + avm + TERM_RESET;
          }
          return []any { avm, tot, " (" + per + ", swap: " + fmt.Sprint(getSwap.Used/1e+6, "/", getSwap.Total/1e+6) + getSwap.Unit + " [" + fmt.Sprint(int(getSwap.PercentageUsed)) + "%])" }
        // }
        // return []any { avm, tot, "" };
      },
    },

    // very fucking intensive
    // {
    //   desc: "cpu",
    //   prependDesc: true,
    //   lnFormat: "%s (cores: %d)",
    //   formatSet: func() []any {
    //     cpu, _ := stats.GetCPU();
    //     var style string;
    //     if strings.Contains(cpu.Model, "Intel") {
    //       style = FORE_CYAN;
    //     } else if strings.Contains(cpu.Model, "AMD") {
    //       style = FORE_RED;
    //     }
    //     return []any { style + cpu.Model + TERM_RESET, cpu.NoOfCores };
    //   },
    // },

    {
      desc: "net",
      prependDesc: true,
      lnFormat: "%s",
      formatSet: func() []any {
        get, _ := stats.GetNetworks();
        if len(get) < 1 {
          return []any { TERM_BLINK_WARN + FORE_RED + "inactive" + TERM_RESET };
        }
        return []any { FORE_GREEN + TERM_BOLD + get[len(get)-1].Interface + TERM_RESET };
      },
    },

    {
      desc: "sto",
      prependDesc: true,
      lnFormat: "%s",
      formatSet: func() []any {
        get, _ := stats.GetDisks();
        if len(get) < 1 {
          return []any {"no disks, what the fuck are you running?"};
        }
        ret := []any {  };
        tmp := "";
        for i := range get {
          if get[i].Type == "efivarfs" || get[i].Type == "vfat" { continue; }
          var use string;
          if get[i].Usage.Available < get[i].Usage.Size / 8 {
            use = FORE_RED;
          }
          tmp += fmt.Sprint(FORE_MAGENTA + TERM_BOLD + get[i].FileSystem + TERM_RESET, " - ", use, int(get[i].Usage.Available)/1.074e+9, "G", TERM_RESET, " left (total ", get[i].Usage.Size/1.074e+9, "G, used ", get[i].Usage.Usage, ", type: ", get[i].Type, ")");
          if get[len(get)-2] != get[i] { tmp += "\n"; }
        }
        ret = append(ret, tmp);
        return []any { tmp };
      },
    },

    {
      desc: "pkg",
      prependDesc: true,
      lnFormat: "%s",
      formatSet: func() []any {
        // filepath.Walk("/var/lib/pacman/local", func(path string, info fs.FileInfo, err error) error {
        //   if info.IsDir() {
        //     pkg++;
        //   }
        //   return nil;
        // });
        f, _ := os.ReadDir("/var/lib/pacman/local");
        pkg := len(f);
        if pkg > 0 { return []any { fmt.Sprint(pkg) + " (" + FORE_YELLOW + "pacman" + TERM_RESET + ")" }; }
        return []any { FORE_RED + "none detected" + TERM_RESET };
      },
    },

    {
      desc: "up",
      prependDesc: true,
      lnFormat: "%s",
      formatSet: func() []any {
        cmd, _ := exec.Command("/usr/bin/uptime", "-p").CombinedOutput();
        if !strings.HasPrefix(string(cmd), "up ") {
          return []any { "err" };
        }
        return []any { strings.TrimSuffix(string(cmd)[3:], "\n") };
      },
    },
  },
};

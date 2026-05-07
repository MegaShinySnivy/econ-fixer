# econ-fixer

> World Fixer for the Space Engineers Economy II update

A utility tool that helps migrate Space Engineers worlds from the pre-Economy II update to the new economy system by automatically removing legacy NPC factions and triggering a fresh regeneration.

---

## What It Does

This program processes your `Sandbox.sbc` save file and:
- **Removes all Economy NPC factions**
- **Preserves player-created factions**
- **Enables auto-generation of new economy stations**

---

## Prerequisites

| Requirement | Description |
|-------------|-------------|
| Old Space Engineers Save | A world that was generated with the pre-Economy II update |
| Binary Download | Get the correct version for your operating system from [Releases](https://github.com/MegaShinySnivy/econ-fixer/releases) |

---

## Quick Start Guide

### 1. Backup Your Save
```
⚠️ Always create a backup of your save directory before proceeding! You should also ensure you actually have old Economy I stations in your world before using this tool. This readme also assumes you know the basics of Space Engineers and its save file structure. 
```

### 2. Prepare Files
Place the following files in the **same directory**:
- `econ-fixer` (the binary executable)
- `Sandbox.sbc` (your Space Engineers world configuration file, case-sensitive!)

### 3. Run The Tool
Execute from your terminal:
```bash
./econ-fixer        # Linux/macOS
.\econ-fixer.exe    # Windows
```

You'll see output listing which factions are being removed or kept:
```
Removing faction: Tag='ABCD' (length=4, meaning it must be an NPC)
Keeping faction: Tag='EFG' (length=3, meaning it must be a player faction)
...
Processing complete. Output written to Sandbox_noNPCs.sbc (XXXX bytes)
```

### 4. Replace The Save File
- Move `Sandbox_noNPCs.sbc` into your Space Engineers save directory
- Rename it to replace the existing `Sandbox.sbc`:
  ```bash
  mv Sandbox_noNPCs.sbc Sandbox.sbc
  # or on Windows:
  move Sandbox_noNPCs.sbc Sandbox.sbc /Y
  ```

### 5. Launch And Verify
1. Start Space Engineers and load your save (the first load may take longer while factions regenerate).
2. Spawn in a space pod or planet rover.
3. Look for a datapad pointing to an Economy station.
4. Ensure the station in question is an enconomy II station. When in doubt, look for the new blocks like the 5x5 connector pad.

---

## Manual Cleanup Required

After running this tool, you'll still need to clean up the following things manually:

### Old Economy I Stations
- All pre-Economy II stations remain in your world, owned by "Unknown".
- Open the Spacemaster menu (`Alt+F10`) and delete the unwanted stations.

### Safe Zones
- Legacy safe zones may still exist around where the old economy stations were.
- Use Spacemaster menu to remove these left over safe zones.

---

## How It Works

The tool uses regex pattern matching to identify and remove faction blocks from the save file:

| Faction Tag | Action | Reason |
|-------------|--------|--------|
| Length = 4 characters | **Removed** | Legacy Economy I NPC factions |
| Longer tags | Kept | Player-created or other factions |
| No Tag field | Kept | Non-faction objects (this should never happen, but better to be safe then risk data loss) |

It then sets `GenerateFactionsOnStart` to `<true>` to ensure new factions spawn when the world is next loaded.

Why use Regex and not the XML parser? Because I am a systems administrator, not a programmer, and I speak Regex far better than I do go libraries.

---

## Troubleshooting

### Issue: "Error reading Sandbox.sbc"
**Cause:** File not found or incorrect path  
**Fix:** Ensure the binary and save file are in the same directory, with exact filename `Sandbox.sbc` (case-sensitive)

### Issue: Nothing changed after running
**Cause:** No NPC factions were detected  
**Fix:** Your world may be missing the Economy system entirely. Good news, you dont need this tool. Just turn it on in the world settings menu before you load the save.

### Issue: New economy stations don't appear
**Cause:** Still using the old save file or failed replacement  
**Fix:** Ensure you replaced `Sandbox.sbc` with `Sandbox_noNPCs.sbc`, not a copy, and that you rename the `Sandbox_noNPCs.sbc` to `Sandbox.sbc` after moving it. I know its a bit confusing, but this is the more safe way to do it. 

### Issue: Game crashes on load
**Cause:** Corrupted save file during processing  
**Fix:** Restore your original backup, and make a bug report with your Sandbox.sbc file before and after editing.

---

## Safety and Fuckups

- **The original `Sandbox.sbc` is never modified**, with the tool creating a new output file instead of overwriting the original.
- Your backup of the entire save directory provides full recovery capability. Don't skip it.
- You can always restore your old Sandbox.sbc if things go wrong, or do a full restore of your save if things go really wrong. Seriously, don't ever skip making a backup.

---

## Requirements Summary

| Platform | Command |
|----------|---------|
| Linux (x86_64) | `./econ-fixer` |
| macOS (Intel/Apple Silicon) | `./econ-fixer` |
| Windows 10+ | `.\econ-fixer.exe` |

---

## Credits

- **Developed by:** MegaShinySnivy  
- **Repository:** [github.com/MegaShinySnivy/econ-fixer](https://github.com/MegaShinySnivy/econ-fixer)  
- **For Space Engineers** (Keen Software House)

---

## License
Licensed under the MIT License. See LICENSE file for details.

---
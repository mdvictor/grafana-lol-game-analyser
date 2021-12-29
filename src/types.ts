import { DataQuery, DataSourceJsonData } from '@grafana/data';

export interface LolCloudDataSourceOptions extends DataSourceJsonData {
  summonerName: string;
  platform: string;
}

export interface LolCloudSecureJsonData {
  apiToken?: string;
}

export interface LolQuery extends DataQuery {
  matchType: string;
  matchId: string;
  timelineData: string;
  player: string;
  normalizeTimerange: boolean;
}

export interface MatchInfo {
  championName: string;
  puuid: string;
  summonerName: string;
  teamId: string;
  individualPosition: string;
  win: boolean;
  kills: number;
  deaths: number;
  assists: number;
  matchId: number;
}

export interface MatchParticipant {
  puuid: string;
  summonerName: string;
  championName: string;
  teamId: number;
  individualPosition: string;
}

export const CHAMPIONS_STATS = 'CHAMPION_STATS';
export const DAMAGE_STATS = 'DAMAGE_STATS';

export const TimelineData = {
  CURRENT_GOLD: 'currentGold',
  TOTAL_GOLD: 'totalGold',
  GOLD_PER_SECOND: 'goldPerSecond',
  MINIONS_KILLED: 'minionsKilled',
  JUNGLE_MINIONS_KILLED: 'jungleMinionsKilled',
  LEVEL: 'level',
  CHAMPION_STATS: {
    ABILITY_HASTE: 'abilityHaste',
    ABILITY_POWER: 'abilityPower',
    ARMOR: 'armor',
    ARMOR_PEN: 'armorPen',
    ATTACK_DAMAGE: 'attackDamage',
    ATTACK_SPEED: 'attackSpeed',
    CC_REDUCTION: 'ccReduction',
    COOLDOWN_REDUCTION: 'cooldownReduction',
    HEALTH_MAX: 'healthMax',
    HEALTH_REGEN: 'healthRegen',
    LIFESTEAL: 'lifesteal',
    MAGIC_PEN: 'magicPen',
    MAGIC_RESIST: 'magicResist',
    MOVEMENT_SPEED: 'movementSpeed',
    OMNIVAMP: 'omnivamp',
    PHYSICAL_VAMP: 'physicalVamp',
    POWER_MAX: 'powerMax',
    SPELL_VAMP: 'spellVamp',
  },
  DAMAGE_STATS: {
    MAGIC_DAMAGE_DONE: 'magicDamageDone',
    MAGIC_DAMAGE_DONE_TO_CHAMPIONS: 'magicDamageDoneToChampions',
    MAGIC_DAMAGE_TAKEN: 'magicDamageTaken',
    PHYSICAL_DAMAGE_DONE: 'physicalDamageDone',
    PHYSICAL_DAMAGE_DONE_TO_CHAMPIONS: 'physicalDamageDoneToChampions',
    PHYSICAL_DAMAGE_TAKEN: 'physicalDamageTaken',
    TOTAL_DAMAGE_DONE: 'totalDamageDone',
    TOTAL_DAMAGE_DONE_TO_CHAMPIONS: 'totalDamageDoneToChampions',
    TOTAL_DAMAGE_TAKEN: 'totalDamageTaken',
    TRUE_DAMAGE_DONE: 'trueDamageDone',
    TRUE_DAMAGE_DONE_TO_CHAMPIONS: 'trueDamageDoneToChampions',
    TRUE_DAMAGE_TAKEN: 'trueDamageTaken',
  },
};

export const RoutingKeys = {
  BR1: 'BR1',
  EUN1: 'EUN1',
  EUW1: 'EUW1',
  JP1: 'JP1',
  KR: 'KR',
  LA1: 'LA1',
  LA2: 'LA2',
  NA1: 'NA1',
  OC1: 'OC1',
  TR1: 'TR1',
  RU: 'RU',
};

export const PlatformRoutingValues = {
  [RoutingKeys.BR1]: 'br1.api.riotgames.com',
  [RoutingKeys.EUN1]: '	eun1.api.riotgames.com',
  [RoutingKeys.EUW1]: 'euw1.api.riotgames.com',
  [RoutingKeys.JP1]: 'jp1.api.riotgames.com',
  [RoutingKeys.KR]: 'kr.api.riotgames.com',
  [RoutingKeys.LA1]: 'la1.api.riotgames.com',
  [RoutingKeys.LA2]: 'la2.api.riotgames.com',
  [RoutingKeys.NA1]: 'na1.api.riotgames.com',
  [RoutingKeys.OC1]: 'oc1.api.riotgames.com',
  [RoutingKeys.TR1]: 'tr1.api.riotgames.com',
  [RoutingKeys.RU]: 'ru.api.riotgames.com',
};

export const MatchTypes = {
  NORMAL: 'normal',
  RANKED: 'ranked',
  TOURNEY: 'tourney',
};
